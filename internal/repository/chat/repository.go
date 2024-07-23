package chat

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/converter"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
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
		if rbErr := tx.Rollback(ctx); err != nil {
			logrus.Errorf("error while rollback transaction: %s", rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

func (r *repository) CreateChat(ctx context.Context, tx pgx.Tx, name string) (int64, error) {
	chatInsert := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_name", "created_at").
		Values(name, time.Now()).
		Suffix("RETURNING id")

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
	chatUserInsert := sq.Insert("chat_users").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "user_name", "created_at").
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
	chatDelete := sq.Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": chatID})

	query, args, err := chatDelete.ToSql()
	if err != nil {
		return err
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *repository) SendMessage(ctx context.Context, tx pgx.Tx, msg *model.Message) error {
	repoMsg := converter.ToRepoFromMessage(msg)

	msgInsert := sq.Insert("chat_messages").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "text", "created_at").
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
