package app

import (
	"context"
	"net"

	"github.com/VadimGossip/consoleChat-chat-server/internal/interceptor"
	"github.com/VadimGossip/consoleChat-chat-server/internal/logger"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	desc "github.com/VadimGossip/consoleChat-chat-server/pkg/chat_v1"
)

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()), grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
		interceptor.ServerTracingInterceptor,
		a.serviceProvider.AuthCheckInterceptor().Hook(),
	)),
	)

	reflection.Register(a.grpcServer)
	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	logger.Infof("[%s] GRPC server is running on: %s", a.name, a.serviceProvider.GRPCConfig().Address())

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
