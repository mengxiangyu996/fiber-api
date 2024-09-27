package config

import "fiber-api/pkg/env"

type RedisConfig struct {
	Host     string
	Port     string
	Database int
	Password string
}

var Redis = &RedisConfig{
	Host:     env.Get("reids.host", "127.0.0.1").(string),
	Port:     env.Get("reids.port", "6379").(string),
	Database: 0,
	Password: env.Get("reids.password", "").(string),
}
