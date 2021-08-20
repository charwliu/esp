package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"go.vixal.xyz/esp/internal/event"
)

var log = event.Log.Sugar()

func logError(result *gorm.DB) {
	if result.Error != nil {
		log.Error(result.Error.Error())
	}
}

type Types map[string]interface{}

// Entities List of database entities and their table names.
var Entities = Types{
	"error":    &Error{},
	"user":     &User{},
	"password": &Password{},
	"dept":     &Dept{},
	"address":  &Address{},
	"menu":     &Menu{},
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
			if err := DB().Raw(fmt.Sprintf("SELECT COUNT(*) AS count FROM %s", name)).Scan(&count).Error; err == nil {
				// log.Debugf("entity: table %s migrated", name)
				break
			} else {
				log.Debugf("entity: wait for migration %s (%s)", err.Error(), name)
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
	for name := range list {
		if err := DB().Exec(fmt.Sprintf("DELETE FROM %s WHERE 1", name)).Error; err == nil {
			// log.Debugf("entity: removed all data from %s", name)
			break
		} else if err.Error() != "record not found" {
			log.Debugf("entity: %s in %s", err, name)
		}
	}
}

// Migrate migrates all database tables of registered entities.
func (list Types) Migrate() {
	for _, entity := range list {
		if err := UnscopedDB().AutoMigrate(entity); err != nil {
			log.Debugf("entity: migrate %s (waiting 1s)", err.Error())

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
	CreateUnknownDept()
	CreateDefaultUsers()

	// CreateUnknownPlace()
	// CreateUnknownLocation()
	// CreateUnknownCountry()
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
		driver = "sqlite3"
		dsn = ".test.db"
	}

	log.Infof("initializing %s test db in %s", driver, dsn)

	db := &Gorm{
		Driver: driver,
		Dsn:    dsn,
	}

	SetDBProvider(db)
	ResetTestFixtures()

	return db
}
