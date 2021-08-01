package database

import (
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

type Database struct {
	*gorm.DB
}

type DefaultModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

func New(config *Config) (*Database, error) {
	var db *gorm.DB
	var err error
	switch strings.ToLower(config.Driver) {
	case "mysql":
		dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.
			Port) + ")/" + config.Database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True"
		db, err = gorm.Open(mysql.Open(dsn),
			//&gorm.Config{},
			&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
		)
		break
	case "postgresql", "postgres":
		dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.
			Database + " host=" + config.Host + " port=" + strconv.Itoa(config.Port)
		db, err = gorm.Open(postgres.Open(dsn),
			//&gorm.Config{},
			&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
		)
		break
	case "sqlserver", "mssql":
		dsn := "sqlserver://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + strconv.Itoa(config.Port) + "?database=" + config.Database
		db, err = gorm.Open(sqlserver.Open(dsn),
			//&gorm.Config{},
			&gorm.Config{Logger: logger.Default.LogMode(logger.Silent)},
		)
		break
	}
	return &Database{db}, err
}
