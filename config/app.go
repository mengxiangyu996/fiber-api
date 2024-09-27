package config

import "fiber-api/pkg/env"

type AppConfig struct {
	Domain string
	Name   string
	Host   string
	Port   int
	Key    string
}

var App = &AppConfig{
	Domain: env.Get("app.domain", "breeze").(string),
	Name:   env.Get("app.name", "breeze").(string),
	Host:   env.Get("app.host", "127.0.0.1").(string),
	Port:   int(env.Get("app.port", 3000).(float64)),
	Key:    env.Get("app.key", "breeze").(string),
}
