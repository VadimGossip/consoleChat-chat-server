package app

import (
	"context"
	"log"

	"github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type DBAdapter struct {
	db       *pgx.Conn
	chatRepo repository.ChatRepository
}

func NewDBAdapter() *DBAdapter {
	return &DBAdapter{}
}

const (
	dbDSN = "host=pg-chat-server port=5432 dbname=chat-server-db user=postgres password=postgres sslmode=disable"
)

func (d *DBAdapter) Connect(ctx context.Context) error {
	db, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err = db.Ping(ctx); err != nil {
		return err
	}
	d.db = db
	d.chatRepo = chat.NewRepository(d.db)

	return nil
}

func (d *DBAdapter) Disconnect(ctx context.Context) error {
	if err := d.db.Close(ctx); err != nil {
		logrus.Errorf("Error occured on db connection close: %s", err.Error())
	}
	return nil
}
