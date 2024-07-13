package main

import (
	"time"

	"github.com/VadimGossip/consoleChat-chat-server/internal/app"
)

var configDir = "config"

func main() {
	auth := app.NewApp("Console Chat Chat-Server", configDir, time.Now())
	auth.Run()
}
