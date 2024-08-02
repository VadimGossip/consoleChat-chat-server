package converter

import (
	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	repoModel "github.com/VadimGossip/consoleChat-chat-server/internal/repository/audit/model"
)

func ToRepoFromAudit(audit *model.Audit) *repoModel.Audit {
	return &repoModel.Audit{
		ID:         audit.ID,
		Action:     audit.Action,
		CallParams: audit.CallParams,
		CreatedAt:  audit.CreatedAt,
	}
}
