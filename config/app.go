package config

import "fiber-api/pkg/env"

type AppConfig struct {
	Host string
	Name string
	Mode string
	Key  string
}

var App = &AppConfig{
	Host: env.Get("app.host", "breeze").(string),
	Name: env.Get("app.name", "breeze").(string),
	Mode: env.Get("app.mode", "debug").(string),
	Key:  env.Get("app.key", "breeze").(string),
}
