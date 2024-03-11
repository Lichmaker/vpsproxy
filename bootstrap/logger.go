package bootstrap

import (
	"github.com/Lichmaker/vpsproxy/pkg/config"
	"github.com/Lichmaker/vpsproxy/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("app.name", "myApp"),
		config.GetString("log.filename", "logs/log.log"),
		config.GetInt("log.max_size", 64),
		config.GetInt("log.max_backup", 5),
		config.GetInt("log.max_age", 30),
		config.GetBool("log.compress", true),
		config.GetString("log.encoder", "console"),
		config.GetString("log.level", "debug"),
	)
}
