package db

import (
	"database/sql"
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Driver       string `yaml:"driver"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	Name         string `yaml:"name"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Locale       string `yaml:"locale"`
	MaxOpenConns int    `yaml:"maxopenconns"`
}

func (d Config) String() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		url.QueryEscape(d.Locale),
	)
}

func Open(cfg Config, debug bool) (*gorm.DB, error) {
	mysqlConn, err := sql.Open("mysql", cfg.String())
	if nil != err {
		return nil, err
	}

	if err = mysqlConn.Ping(); nil != err {
		return nil, err
	}

	mysqlConn.SetMaxIdleConns(0)
	mysqlConn.SetMaxOpenConns(cfg.MaxOpenConns)

	dbConn, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DriverName: "mysql",
				Conn:       mysqlConn,
			},
		),
	)

	logLevel := logger.Silent
	if debug {
		logLevel = logger.Info
	}

	dbConn.Logger.LogMode(logLevel)

	return dbConn, nil
}
