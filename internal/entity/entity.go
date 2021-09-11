package entity

import (
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"go.vixal.xyz/esp/internal/event"
)

func S() *zap.SugaredLogger {
	return event.S()
}

func L() *zap.Logger {
	return event.L()
}

type Types map[string]interface{}

// Entities List of database entities and their table names.
var Entities = Types{
	"error":      &Error{},
	"address":    &Address{},
	"user":       &User{},
	"department": &Department{},
	"password":   &Password{},
	"menu":       &Menu{},
	"country":    &Country{},
	"place":      &Place{},
	"cell":       &Cell{},
	"label":      &Label{},
	"link":       &Link{},
	"categories": &Category{},
	// "organization":              &Organization{},
	// "oauth_client":              &OauthClient{},
	// "user_registration_profile": &UserRegistrationProfile{},
	// "auth_token":                &AuthToken{},
	// "pep_proxy":                 &PepProxy{},

}

type RowCount struct {
	Count int
}

// WaitForMigration waits for the database migration to be successful.
func (list Types) WaitForMigration() {
	attempts := 100
	for name := range list {
		for i := 0; i <= attempts; i++ {
			count := RowCount{}
			if err := DB().Raw(fmt.Sprintf(`SELECT COUNT(*) AS count FROM "%s"`, name)).Scan(&count).Error; err == nil {
				// log.Debug("entity: table migrated", zap.String("name", name))
				break
			} else {
				L().Debug("entity: wait for migration",
					zap.String("name", name),
					zap.Error(err))
			}

			if i == attempts {
				panic("migration failed")
			}

			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Truncate removes all data from tables without dropping them.
func (list Types) Truncate() {
	var retries []string
	for name := range list {
		if err := DB().Exec(fmt.Sprintf(`DELETE FROM "%s" WHERE 1 = 1`, name)).Error; err == nil {
			L().Debug("entity: removed all data from", zap.String("name", name))
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			L().Debug("entity:", zap.String("name", name), zap.Error(err))
		} else {
			retries = append(retries, name)
		}
	}
	for _, name := range retries {
		if err := DB().Exec(fmt.Sprintf(`DELETE FROM "%s" WHERE 1 = 1`, name)).Error; err == nil {
			L().Debug("entity: removed all data from", zap.String("name", name))
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			L().Debug("entity:", zap.String("name", name), zap.Error(err))
		}
	}
}

// Migrate migrates all database tables of registered entities.
func (list Types) Migrate() {
	for _, entity := range list {
		if err := UnscopedDB().AutoMigrate(entity); err != nil {
			L().Debug("entity: migrate (waiting 1s)", zap.Error(err))

			time.Sleep(time.Second)

			if err := UnscopedDB().AutoMigrate(entity); err != nil {
				panic(err)
			}
		}
	}
}

// Drop drops all database tables of registered entities.
func (list Types) Drop() {
	for _, entity := range list {
		if err := UnscopedDB().Migrator().DropTable(entity); err != nil {
			panic(err)
		}
	}
}

// CreateDefaultFixtures Creates default database entries for test and production.
func CreateDefaultFixtures() {
	CreateUnknownAddress()
	CreateUnknownDepartment()
	CreateDefaultUsers()

	CreateUnknownPlace()
	CreateUnknownLocation()
	CreateUnknownCountry()
	CreateUnknownMenu()
	// CreateUnknownCamera()
	// CreateUnknownLens()
}

// MigrateDb creates all tables and inserts default entities as needed.
func MigrateDb() {
	Entities.Migrate()
	Entities.WaitForMigration()

	CreateDefaultFixtures()
}

// ResetTestFixtures drops database tables for all known entities and re-creates them with fixtures.
func ResetTestFixtures() {
	Entities.Migrate()
	Entities.WaitForMigration()
	Entities.Truncate()

	CreateDefaultFixtures()

	CreateTestFixtures()
}

// InitTestDb connects to and completely initializes the test database incl fixtures.
func InitTestDb(driver, dsn string) *Gorm {
	if HasDBProvider() {
		return nil
	}

	if driver == "test" || driver == "sqlite" || driver == "" || dsn == "" {
		driver = SQLite
		dsn = ".test.db"
	}
	if driver == "postgresql" {
		driver = Postgres
	}

	L().Info("initializing test db",
		zap.String("driver", driver),
		zap.String("dsn", dsn))

	db := &Gorm{
		Driver: driver,
		Dsn:    dsn,
	}
	if err := db.Connect(); err != nil {
		L().Fatal(err.Error())
	}
	SetDBProvider(db)
	ResetTestFixtures()

	return db
}
