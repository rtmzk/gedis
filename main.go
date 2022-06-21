package main

import (
	"fmt"
	"github.com/rtmzk/gedis/internal/handlers"
	"github.com/rtmzk/gedis/internal/server"
	"github.com/rtmzk/gedis/pkg/config"
	"github.com/rtmzk/gedis/pkg/logger"
	"os"
)

const configFile string = "redis.conf"

var defaultProperties = &config.ServerProperties{
	Bind: "0.0.0.0",
	Port: 6379,
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func main() {
	logger.Setup(&logger.Settings{
		Path:       "logs",
		Name:       "gedis",
		Ext:        "log",
		TimeFormat: "2006-01-02",
	})
	if fileExists(configFile) {
		config.SetupConfig(configFile)
	} else {
		config.Properties = defaultProperties
	}
	err := server.ListenAndServeWithSignal(&server.Config{
		Address: fmt.Sprintf("%s:%d", config.Properties.Bind, config.Properties.Port),
	}, handlers.NewEchoHandler())

	if err != nil {
		logger.Error(err)
	}
}
