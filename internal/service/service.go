package service

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, chatID int64) error
	SendMessage(ctx context.Context, msg *model.Message) error
}

type AuditService interface {
	Create(ctx context.Context, audit *model.Audit) error
}
