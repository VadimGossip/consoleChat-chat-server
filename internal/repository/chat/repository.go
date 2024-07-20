package chat

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"sync"
	"time"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/converter"
	repoModel "github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/model"
	"github.com/sirupsen/logrus"
)

var _ def.ChatRepository = (*repository)(nil)

type repository struct {
	db     *pgx.Conn
	m      sync.RWMutex
	data   map[int64]*repoModel.Chat
	lastID int64
}

func NewRepository(db *pgx.Conn) *repository {
	return &repository{
		db:   db,
		data: make(map[int64]*repoModel.Chat),
	}
}

func (r *repository) Create(_ context.Context, usernames []string) (int64, error) {
	r.m.Lock()
	defer r.m.Unlock()

	r.lastID++
	r.data[r.lastID] = &repoModel.Chat{
		ID:        r.lastID,
		Users:     usernames,
		CreatedAt: time.Now(),
		Messages:  make([]repoModel.Message, 0),
	}
	logrus.Infof("Chat id %d for users %v", r.lastID, r.data[r.lastID])
	return r.lastID, nil
}

func (r *repository) Delete(_ context.Context, id int64) error {
	r.m.Lock()
	defer r.m.Unlock()
	if _, ok := r.data[id]; ok {
		delete(r.data, id)
		logrus.Infof("Chat id=%d deleted", id)
		return nil
	}
	return fmt.Errorf("chat id=%d not found", id)
}

func (r *repository) SendMessage(_ context.Context, id int64, msg *model.Message) error {
	r.m.Lock()
	defer r.m.Unlock()
	if chat, ok := r.data[id]; ok {
		chat.Messages = append(chat.Messages, converter.ToRepoFromMessage(msg))
		logrus.Infof("Message added %+v", *msg)
		logrus.Infof("All msgs %+v", chat.Messages)
		return nil
	}
	return fmt.Errorf("chat id=%d not found", id)
}
