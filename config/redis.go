package config

import "fiber-api/pkg/env"

type RedisConfig struct {
	Host     string
	Port     int
	Database int
	Password string
}

var Redis = &RedisConfig{
	Host:     env.Get("reids.host", "127.0.0.1").(string),
	Port:     int(env.Get("reids.port", 6379).(float64)),
	Database: int(env.Get("reids.database", 0).(float64)),
	Password: env.Get("reids.password", "").(string),
}
