package chat

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/converter"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const (
	chats        string = "chats"
	columnId     string = "id"
	chatName     string = "chat_name"
	createdAt    string = "created_at"
	chatUsers    string = "chat_users"
	chatId       string = "chat_id"
	userId       string = "user_id"
	username     string = "username"
	chatMessages string = "chat_messages"
	text         string = "text"
)

var _ def.ChatRepository = (*repository)(nil)

type repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) BeginTxSerializable(ctx context.Context) (pgx.Tx, error) {
	return r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
}

func (r *repository) StopTx(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			logrus.Errorf("error while rollback transaction: %s", rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

func (r *repository) CreateChat(ctx context.Context, tx pgx.Tx, name string) (int64, error) {
	chatInsert := sq.Insert(chats).
		PlaceholderFormat(sq.Dollar).
		Columns(chatName, createdAt).
		Values(name, time.Now()).
		Suffix("RETURNING " + columnId)

	query, args, err := chatInsert.ToSql()
	if err != nil {
		return 0, err
	}
	var id int64
	if err = tx.QueryRow(ctx, query, args...).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) CreateChatUser(ctx context.Context, tx pgx.Tx, chatID int64, user model.User) error {
	chatUserInsert := sq.Insert(chatUsers).
		PlaceholderFormat(sq.Dollar).
		Columns(chatId, userId, username, createdAt).
		Values(chatID, user.ID, user.Name, time.Now())

	query, args, err := chatUserInsert.ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, tx pgx.Tx, chatID int64) error {
	chatDelete := sq.Delete(chats).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnId: chatID})

	query, args, err := chatDelete.ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *repository) SendMessage(ctx context.Context, tx pgx.Tx, msg *model.Message) error {
	repoMsg := converter.ToRepoFromMessage(msg)

	msgInsert := sq.Insert(chatMessages).
		PlaceholderFormat(sq.Dollar).
		Columns(chatId, userId, text, createdAt).
		Values(repoMsg.ChatID, repoMsg.UserID, repoMsg.Text, repoMsg.CreatedAt)

	query, args, err := msgInsert.ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}
