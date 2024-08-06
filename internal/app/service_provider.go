package app

import (
	"context"
	"fmt"
	"log"

	"github.com/VadimGossip/consoleChat-chat-server/internal/api/chat"
	"github.com/VadimGossip/consoleChat-chat-server/internal/client/db"
	"github.com/VadimGossip/consoleChat-chat-server/internal/client/db/pg"
	"github.com/VadimGossip/consoleChat-chat-server/internal/client/db/transaction"
	"github.com/VadimGossip/consoleChat-chat-server/internal/closer"
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	auditRepo "github.com/VadimGossip/consoleChat-chat-server/internal/repository/audit"
	chatRepo "github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service"
	auditService "github.com/VadimGossip/consoleChat-chat-server/internal/service/audit"
	chatService "github.com/VadimGossip/consoleChat-chat-server/internal/service/chat"
	"github.com/sirupsen/logrus"
)

type serviceProvider struct {
	cfg *model.Config

	dbClient  db.Client
	txManager db.TxManager
	auditRepo repository.AuditRepository
	chatRepo  repository.ChatRepository

	auditService service.AuditService
	chatService  service.ChatService

	chatImpl *chat.Implementation
}

func newServiceProvider(cfg *model.Config) *serviceProvider {
	return &serviceProvider{cfg: cfg}
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		dbDSN := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", s.cfg.Db.Host, s.cfg.Db.Port, s.cfg.Db.Name, s.cfg.Db.Username, s.cfg.Db.Password, s.cfg.Db.SSLMode)
		cl, err := pg.New(ctx, dbDSN)
		if err != nil {
			logrus.Fatalf("failed to create db client: %s", err)
		}

		if err = cl.DB().Ping(ctx); err != nil {
			log.Fatalf("ping error: %s", err)
		}
		closer.Add(cl.Close)
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AuditRepository(ctx context.Context) repository.AuditRepository {
	if s.auditRepo == nil {
		s.auditRepo = auditRepo.NewRepository(s.DBClient(ctx))
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
		s.chatRepo = chatRepo.NewRepository(s.DBClient(ctx))
	}
	return s.chatRepo
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.chatRepository(ctx), s.AuditService(ctx), s.TxManager(ctx))
	}
	return s.chatService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
