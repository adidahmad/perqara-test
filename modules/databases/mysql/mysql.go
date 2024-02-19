package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/adidahmad/perqara-test/config"
	_ "github.com/go-sql-driver/mysql"
)

func MySQLConnect(config config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.MaxLifetime * time.Second)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
