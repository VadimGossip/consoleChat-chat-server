package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/VadimGossip/consoleChat-chat-server/internal/app"
)

var configDir = "config"
var appName = "Console Chat Chat-Server"

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx, appName, configDir, time.Now())
	if err != nil {
		logrus.Fatalf("failed to init app[%s]: %s", appName, err)
	}

	if err = a.Run(); err != nil {
		logrus.Infof("Application run process finished with error: %s", err)
	}
}
