package chat

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	//"github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/converter"
	//repoModel "github.com/VadimGossip/consoleChat-chat-server/internal/repository/chat/model"
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
		if err := tx.Rollback(ctx); err != nil {
			logrus.Errorf("error while rollback transaction: %s", err)
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

func (r *repository) CreateChatUser(ctx context.Context, tx pgx.Tx, chatId int64, user model.User) error {
	chatUsersInsert := sq.Insert("chat_users").
		PlaceholderFormat(sq.Dollar).
		Columns("chat_id", "user_id", "user_name", "created_at").
		Values(chatId, user.Id, user.Name, time.Now())

	query, args, err := chatUsersInsert.ToSql()
	if err != nil {
		return err
	}
	fmt.Println(query, args)

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(_ context.Context, id int64) error {

	return fmt.Errorf("chat id=%d not found", id)
}

func (r *repository) SendMessage(_ context.Context, id int64, msg *model.Message) error {
	return fmt.Errorf("chat id=%d not found", id)
}
