package config

import "github.com/Lichmaker/vpsproxy/pkg/config"

func init() {
	config.BindEnv("log.level", "LOG_LEVEL")
	config.BindEnv("log.encoder", "LOG_ENCODER")
	config.BindEnv("log.filename", "LOG_NAME")
	config.BindEnv("log.max_size", "LOG_MAX_SIZE")
	config.BindEnv("log.max_backup", "LOG_MAX_BACKUP")
	config.BindEnv("log.max_age", "LOG_MAX_AGE")
	config.BindEnv("log.compress", "LOG_COMPRESS")
}
