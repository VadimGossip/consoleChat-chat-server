package main

import (
	"time"

	"github.com/VadimGossip/consoleChat-chat-server/internal/app"
	"github.com/sirupsen/logrus"
)

var configDir = "config"

func main() {
	auth := app.NewApp("Console Chat Chat-Server", configDir, time.Now())
	if err := auth.Run(); err != nil {
		logrus.Infof("Application run process finished with error")
	}
}
