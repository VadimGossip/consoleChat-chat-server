package repository

import (
	"context"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	"github.com/jackc/pgx/v4"
)

type ChatRepository interface {
	BeginTxSerializable(ctx context.Context) (pgx.Tx, error)
	StopTx(ctx context.Context, tx pgx.Tx, err error) error
	CreateChat(ctx context.Context, tx pgx.Tx, name string) (int64, error)
	CreateChatUser(ctx context.Context, tx pgx.Tx, chatID int64, user model.User) error
	Delete(ctx context.Context, tx pgx.Tx, chatID int64) error
	SendMessage(ctx context.Context, tx pgx.Tx, msg *model.Message) error
}
