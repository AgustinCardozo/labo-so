package main

import (
	"fmt"
	"github.com/AgustinCardozo/labo-so/cpu/models/global"
	"github.com/AgustinCardozo/labo-so/utils/config"
	"github.com/AgustinCardozo/labo-so/utils/greetings"
	"github.com/AgustinCardozo/labo-so/utils/log"
	"log/slog"
)

const (
	ConfigPath = "./config/cpu.json"
	LogPath    = "cpu.log"
)

func main() {
	message := greetings.Hello("Burak")
	fmt.Println(message)

	config.InitConfig(ConfigPath, &global.ConfigCpu)
	log.InitLogger(LogPath, global.ConfigCpu.LogLevel)

	slog.Debug("test")
}
