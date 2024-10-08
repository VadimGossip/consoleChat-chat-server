package chat

import (
	db "github.com/VadimGossip/platform_common/pkg/db/postgres"

	"github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/service"
)

var _ def.ChatService = (*service)(nil)

type service struct {
	chatRepository repository.ChatRepository
	auditService   def.AuditService
	txManager      db.TxManager
}

func NewService(chatRepository repository.ChatRepository, auditService def.AuditService, txManager db.TxManager) *service {
	return &service{
		chatRepository: chatRepository,
		auditService:   auditService,
		txManager:      txManager,
	}
}
