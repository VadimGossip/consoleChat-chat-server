package repository

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
)

type ChatRepository interface {
	CreateChat(ctx context.Context, name string) (int64, error)
	CreateChatUser(ctx context.Context, chatID int64, user model.User) error
	Delete(ctx context.Context, chatID int64) error
	SendMessage(ctx context.Context, msg *model.Message) error
}

type AuditRepository interface {
	Create(ctx context.Context, audit *model.Audit) error
}
