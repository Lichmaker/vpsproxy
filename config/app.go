package config

import "github.com/Lichmaker/vpsproxy/pkg/config"

func init() {
	config.BindEnv("app.host", "APP_HOST")
	config.BindEnv("app.port", "APP_PORT")
	config.BindEnv("app.name", "APP_NAME")
	config.BindEnv("app.env", "APP_ENV")
	config.BindEnv("app.debug", "APP_DEBUG")
	config.BindEnv("app.timezone", "TIMEZONE")
}
