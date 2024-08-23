package app

import (
	"context"
	"log"

	"github.com/VadimGossip/platform_common/pkg/closer"
	"github.com/VadimGossip/platform_common/pkg/db/postgres"
	"github.com/VadimGossip/platform_common/pkg/db/postgres/pg"
	"github.com/VadimGossip/platform_common/pkg/db/postgres/transaction"
	"github.com/sirupsen/logrus"

	"github.com/VadimGossip/consoleChat-chat-server/internal/api/chat"
	"github.com/VadimGossip/consoleChat-chat-server/internal/client/grpc"
	"github.com/VadimGossip/consoleChat-chat-server/internal/client/grpc/auth"
	"github.com/VadimGossip/consoleChat-chat-server/internal/config"
	clientCfg "github.com/VadimGossip/consoleChat-chat-server/internal/config/client"
	dbCfg "github.com/VadimGossip/consoleChat-chat-server/internal/config/db"
	serverCfg "github.com/VadimGossip/consoleChat-chat-server/internal/config/server"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	auditRepo "github.com/VadimGossip/consoleChat-chat-server/internal/repository/audit"
	chatRepo "github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service"
	auditService "github.com/VadimGossip/consoleChat-chat-server/internal/service/audit"
	chatService "github.com/VadimGossip/consoleChat-chat-server/internal/service/chat"
)

type serviceProvider struct {
	grpcConfig           config.GRPCConfig
	pgConfig             config.PgConfig
	authGRPCClientConfig config.AuthGRPCClientConfig

	pgDbClient postgres.Client
	txManager  postgres.TxManager
	auditRepo  repository.AuditRepository
	chatRepo   repository.ChatRepository

	auditService service.AuditService
	chatService  service.ChatService

	authGRPCClient grpc.AuthClient

	chatImpl *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := serverCfg.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpcConfig: %s", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PGConfig() config.PgConfig {
	if s.pgConfig == nil {
		cfg, err := dbCfg.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pgConfig: %s", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) AuthGRPCClientConfig() config.AuthGRPCClientConfig {
	if s.authGRPCClientConfig == nil {
		cfg, err := clientCfg.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get authGRPCClientConfig: %s", err)
		}

		s.authGRPCClientConfig = cfg
	}

	return s.authGRPCClientConfig
}

func (s *serviceProvider) PgDbClient(ctx context.Context) postgres.Client {
	if s.pgDbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			logrus.Fatalf("failed to create db client: %s", err)
		}

		if err = cl.DB().Ping(ctx); err != nil {
			log.Fatalf("ping error: %s", err)
		}
		closer.Add(cl.Close)
		s.pgDbClient = cl
	}

	return s.pgDbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) postgres.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.PgDbClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AuditRepository(ctx context.Context) repository.AuditRepository {
	if s.auditRepo == nil {
		s.auditRepo = auditRepo.NewRepository(s.PgDbClient(ctx))
	}
	return s.auditRepo
}

func (s *serviceProvider) AuditService(ctx context.Context) service.AuditService {
	if s.auditService == nil {
		s.auditService = auditService.NewService(s.AuditRepository(ctx))
	}
	return s.auditService
}

func (s *serviceProvider) chatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepo == nil {
		s.chatRepo = chatRepo.NewRepository(s.PgDbClient(ctx))
	}
	return s.chatRepo
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.chatRepository(ctx), s.AuditService(ctx), s.TxManager(ctx))
	}
	return s.chatService
}

func (s *serviceProvider) AuthGRPCClient() grpc.AuthClient {
	if s.authGRPCClient == nil {
		grpcAuthClient, err := auth.NewClient(s.AuthGRPCClientConfig())
		if err != nil {
			logrus.Fatalf("failed to create access grpc client: %s", err)
		}
		s.authGRPCClient = grpcAuthClient
	}
	return s.authGRPCClient
}

func (s *serviceProvider) UserImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
