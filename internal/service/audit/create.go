package audit

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
)

func (s *service) Create(ctx context.Context, audit *model.Audit) error {
	return s.auditRepository.Create(ctx, audit)
}
