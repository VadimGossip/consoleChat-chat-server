package app

import (
	"context"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat"
	"log"

	"github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	"github.com/jackc/pgx/v4"
)

type DBAdapter struct {
	chatRepo repository.ChatRepository
}

func NewDBAdapter() *DBAdapter {
	return &DBAdapter{}
}

const (
	dbDSN = "host=localhost port=54321 dbname=chat-server-db user=postgres password=postgres sslmode=disable"
)

func (d *DBAdapter) Connect(ctx context.Context) error {
	db, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close(ctx)

	if err = db.Ping(ctx); err != nil {
		return err
	}
	d.chatRepo = chat.NewRepository(db)

	return nil
}
