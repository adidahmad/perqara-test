package config

import (
	"time"

	"github.com/spf13/viper"
)

type Database struct {
	Host         string
	User         string
	Password     string
	DBName       string
	Port         int
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  time.Duration
}

func loadDBConfig(name string) {
	db := viper.Sub("database." + name)
	AppConf.DBConfig = Database{
		Host:         db.GetString("host"),
		User:         db.GetString("user"),
		Password:     db.GetString("password"),
		DBName:       db.GetString("db_name"),
		Port:         db.GetInt("port"),
		MaxOpenConns: db.GetInt("pool.max_open_conns"),
		MaxIdleConns: db.GetInt("pool.max_idle_conns"),
		MaxLifetime:  db.GetDuration("pool.max_lifetime"),
	}
}
