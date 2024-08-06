package audit

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/service"
)

var _ def.AuditService = (*service)(nil)

type service struct {
	auditRepository repository.AuditRepository
}

func NewService(auditRepository repository.AuditRepository) *service {
	return &service{
		auditRepository: auditRepository,
	}
}
