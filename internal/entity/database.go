package entity

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"go.vixal.xyz/esp/internal/mutex"
)

// Database drivers (sql dialects).
const (
	MySQL     = "mysql"
	MariaDB   = "mariadb"
	SQLite    = "sqlite3"
	Postgres  = "postgres"
	SQLServer = "sqlserver"
)

type DBConfig struct {
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
	if g.db == nil {
		L().Fatal("entity: database not connected")
	}
	return g.db
}

func (g *Gorm) Connect() error {
	mutex.DB.Lock()
	defer mutex.DB.Unlock()
	dbDriver := g.Driver
	dbDsn := g.Dsn

	if dbDriver == "" {
		return errors.New("config: database driver not specified")
	}

	if dbDsn == "" {
		return errors.New("config: database DSN not specified")
	}

	var db *gorm.DB
	var err error
	var i int
	for {
		switch dbDriver {
		case MySQL, MariaDB:
			db, err = gorm.Open(mysql.Open(dbDsn),
				&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			)
			break
		case Postgres:
			db, err = gorm.Open(postgres.Open(dbDsn),
				&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
			)
			break
		case SQLite:
			db, err = gorm.Open(sqlite.Open(dbDsn),
				&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
			)
			break
		}

		if i == 12 || (db != nil && err == nil) {
			break
		} else {
			i += 1
			time.Sleep(5 * time.Second)
		}
	}
	if err != nil || db == nil {
		L().Fatal(err.Error())
	}

	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(4)
		sqlDB.SetMaxOpenConns(256)
		sqlDB.SetConnMaxLifetime(10 * time.Minute)
	}
	g.db = db
	return err
}

type DefaultModel struct {
	ID        *uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (base *DefaultModel) BeforeCreate(*gorm.DB) error {
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
			L().Fatal(err.Error())
		}
		g.db = nil
	}
}
