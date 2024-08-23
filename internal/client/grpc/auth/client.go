package auth

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	descAccess "github.com/VadimGossip/concoleChat-auth/pkg/access_v1"
	descGrpc "github.com/VadimGossip/consoleChat-chat-server/internal/client/grpc"
	"github.com/VadimGossip/consoleChat-chat-server/internal/config"
)

type client struct {
	cl descAccess.AccessV1Client
}

func NewClient(authGRPCClientConfig config.AuthGRPCClientConfig) (descGrpc.AuthClient, error) {
	conn, err := grpc.NewClient(authGRPCClientConfig.Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Errorf("failed to connect to grpc server: %v", err)
	}
	return &client{cl: descAccess.NewAccessV1Client(conn)}, nil
}

func (c *client) Check(ctx context.Context, endpointAddress string) error {
	_, err := c.cl.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: endpointAddress,
	})
	if err != nil {
		return err
	}
	return nil
}
