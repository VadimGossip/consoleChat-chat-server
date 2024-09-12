package client

import (
	"fmt"
	"os"
	"strconv"

	"github.com/VadimGossip/consoleChat-chat-server/internal/logger"
	"github.com/pkg/errors"
)

const (
	authGRPCHostEnvName = "AUTH_GRPC_SERVER_HOST"
	authGRPCPortEnvName = "AUTH_GRPC_SERVER_PORT"
)

type authGRPCClientConfig struct {
	host string
	port int
}

func (cfg *authGRPCClientConfig) setFromEnv() error {
	var err error
	cfg.host = os.Getenv(authGRPCHostEnvName)
	if len(cfg.host) == 0 {
		return fmt.Errorf("authGRPCClientConfig host not found")
	}

	portStr := os.Getenv(authGRPCPortEnvName)
	if len(portStr) == 0 {
		return fmt.Errorf("authGRPCClientConfig port not found")
	}

	cfg.port, err = strconv.Atoi(portStr)
	if err != nil {
		return errors.Wrap(err, "failed to parse authGRPCClientConfig port")
	}
	return nil
}

func NewGRPCConfig() (*authGRPCClientConfig, error) {
	cfg := &authGRPCClientConfig{}
	if err := cfg.setFromEnv(); err != nil {
		return nil, fmt.Errorf("authGRPCClientConfig set from env err: %s", err)
	}

	logger.Infof("authGRPCClientConfig: [%+v]", *cfg)
	return cfg, nil
}

func (cfg *authGRPCClientConfig) Address() string {
	return fmt.Sprintf("%s:%d", cfg.host, cfg.port)
}
