package repository

import (
	"context"
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, usernames []string) (int64, error)
	Delete(ctx context.Context, id int64) error
	SendMessage(ctx context.Context, id int64, msg *model.Message) error
}
