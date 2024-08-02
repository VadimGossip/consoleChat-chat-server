package chat

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/VadimGossip/consoleChat-chat-server/internal/client/db"
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/converter"
)

const (
	chatsTableName        string = "chats"
	chatUsersTableName    string = "chat_users"
	chatMessagesTableName string = "chat_messages"
	idColumn              string = "id"
	chatNameColumn        string = "chat_name"
	createdAtColumn       string = "created_at"
	chatIDColumn          string = "chat_id"
	userIDColumn          string = "user_id"
	usernameColumn        string = "username"
	textColumn            string = "text"
	repoName              string = "chat_repository"
)

var _ def.ChatRepository = (*repository)(nil)

type repository struct {
	db db.Client
}

func NewRepository(db db.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateChat(ctx context.Context, name string) (int64, error) {
	chatInsert := sq.Insert(chatsTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatNameColumn, createdAtColumn).
		Values(name, time.Now()).
		Suffix("RETURNING " + idColumn)

	query, args, err := chatInsert.ToSql()
	if err != nil {
		return 0, err
	}
	var id int64
	q := db.Query{
		Name:     repoName + ".CreateChat",
		QueryRaw: query,
	}
	if err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) CreateChatUser(ctx context.Context, chatID int64, user model.User) error {
	repoUser := converter.ToRepoFromUser(user)
	chatUserInsert := sq.Insert(chatUsersTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userIDColumn, usernameColumn, createdAtColumn).
		Values(chatID, repoUser.ID, repoUser.Name, time.Now())

	query, args, err := chatUserInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     repoName + ".CreateChatUser",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, chatID int64) error {
	chatDelete := sq.Delete(chatsTableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: chatID})

	query, args, err := chatDelete.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     repoName + ".Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) SendMessage(ctx context.Context, msg *model.Message) error {
	repoMsg := converter.ToRepoFromMessage(msg)

	msgInsert := sq.Insert(chatMessagesTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(chatIDColumn, userIDColumn, textColumn, createdAtColumn).
		Values(repoMsg.ChatID, repoMsg.UserID, repoMsg.Text, repoMsg.CreatedAt)

	query, args, err := msgInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     repoName + ".SendMessage",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
