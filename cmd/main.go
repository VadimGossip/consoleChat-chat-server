package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/VadimGossip/consoleChat-chat-server/internal/logger"
	"github.com/VadimGossip/consoleChat-chat-server/internal/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/VadimGossip/consoleChat-chat-server/internal/app"
)

var appName = "Console Chat Chat-Server"
var logLevel = flag.String("l", "info", "log level")

func main() {
	ctx := context.Background()
	flag.Parse()
	logger.Init(getCore(getAtomicLevel(*logLevel)))
	tracing.Init(logger.Logger(), appName)

	a, err := app.NewApp(ctx, appName, time.Now())
	if err != nil {
		logger.Fatalf("failed to init app[%s]: %s", appName, err)
	}

	if err = a.Run(); err != nil {
		logger.Infof("Application run process finished with error: %s", err)
	}
}

func getCore(level zap.AtomicLevel) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)

	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	return zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
	)
}

func getAtomicLevel(loglevel string) zap.AtomicLevel {
	var level zapcore.Level
	if err := level.Set(loglevel); err != nil {
		log.Fatalf("failed to set log level: %v", err)
	}

	return zap.NewAtomicLevelAt(level)
}
