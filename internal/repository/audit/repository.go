package audit

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/VadimGossip/consoleChat-chat-server/internal/client/db"
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	def "github.com/VadimGossip/consoleChat-chat-server/internal/repository"
	"github.com/VadimGossip/consoleChat-chat-server/internal/repository/audit/converter"
)

const (
	auditTableName   string = "audit_log"
	actionColumn     string = "action"
	callParamsColumn string = "call_params"
	createdAtColumn  string = "created_at"
	repoName         string = "audit_repository"
)

var _ def.AuditRepository = (*repository)(nil)

type repository struct {
	db db.Client
}

func NewRepository(db db.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, audit *model.Audit) error {
	repoAudit := converter.ToRepoFromAudit(audit)
	userInsert := sq.Insert(auditTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(actionColumn, callParamsColumn, createdAtColumn).
		Values(repoAudit.Action, repoAudit.CallParams, time.Now())

	query, args, err := userInsert.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     repoName + ".Create",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
