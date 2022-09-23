package db

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
