package entity

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database drivers (sql dialects).
const (
	MySQL    = "mysql"
	SQLite   = "sqlite3"
	Postgres = "postgres"
)

type Config struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

type DBProvider interface {
	DB() *gorm.DB
}

var dbProvider DBProvider

func IsDBDialect(name string) bool {
	return name == DB().Dialector.Name()
}

func DBDialect() string {
	return DB().Dialector.Name()
}

func SetDBProvider(provider DBProvider) {
	dbProvider = provider
}

func HasDBProvider() bool {
	return dbProvider != nil
}

func UnscopedDB() *gorm.DB {
	return DB().Unscoped()

}

func DB() *gorm.DB {
	return dbProvider.DB()
}

type Gorm struct {
	Driver string
	Dsn    string

	once sync.Once
	db   *gorm.DB
}

func (g *Gorm) DB() *gorm.DB {
	g.once.Do(g.Connect)
	if g.db == nil {
		log.Fatal("entity: database not connected")
	}
	return g.db
}

func (g *Gorm) Connect() {
	var db *gorm.DB
	var err error
	var i int
	for {
		switch g.Driver {
		case MySQL:
			db, err = gorm.Open(mysql.Open(g.Dsn),
				&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			)
			break
		case Postgres:
			db, err = gorm.Open(postgres.Open(g.Dsn),
				&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			)
			break
		case SQLite:
			db, err = gorm.Open(sqlite.Open(g.Dsn),
				&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			)
			break // Not required as unicode is default.
		}
		if i == 12 || (db != nil && err == nil) {
			break
		} else {
			i += 1
			time.Sleep(5 * time.Second)
		}
	}
	if err != nil || db == nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(4)
		sqlDB.SetMaxOpenConns(256)
	}
	g.db = db
}

type DefaultModel struct {
	ID        *uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

func (base *DefaultModel) BeforeCreate(scope *gorm.DB) error {
	if base.ID == nil {
		id := uuid.New()
		base.ID = &id
	}
	return nil
}

func (g *Gorm) Close() {
	if g.db == nil {
		return
	}
	if sqlDB, e := g.db.DB(); e == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatal(err)
		}
		g.db = nil
	}
}
