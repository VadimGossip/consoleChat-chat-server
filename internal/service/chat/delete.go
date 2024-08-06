package chat

import (
	"context"
	"fmt"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
)

func (s *service) Delete(ctx context.Context, chatID int64) error {
	return s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		if txErr = s.chatRepository.Delete(ctx, chatID); txErr != nil {
			return txErr
		}

		txErr = s.auditService.Create(ctx, &model.Audit{
			Action:     "delete chat",
			CallParams: fmt.Sprintf("chatID %d", chatID),
		})
		if txErr != nil {
			return txErr
		}

		return nil
	})
}
