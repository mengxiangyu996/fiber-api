package config

import "fiber-api/pkg/env"

type MysqlConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

var Mysql = &MysqlConfig{
	Host:     env.Get("mysql.host", "127.0.0.1").(string),
	Port:     env.Get("mysql.port", "3306").(string),
	Database: env.Get("mysql.database", "fiber-api").(string),
	Username: env.Get("mysql.username", "root").(string),
	Password: env.Get("mysql.password", "root").(string),
}
