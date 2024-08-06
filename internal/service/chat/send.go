package chat

import (
	"context"
	"fmt"

	"github.com/VadimGossip/consoleChat-chat-server/internal/model"
	"github.com/VadimGossip/consoleChat-chat-server/internal/service/chat/validator"
)

func (s *service) SendMessage(ctx context.Context, msg *model.Message) error {
	if err := validator.SendValidation(msg); err != nil {
		return err
	}
	return s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var txErr error
		if txErr = s.chatRepository.SendMessage(ctx, msg); txErr != nil {
			return txErr
		}

		if txErr = s.auditService.Create(ctx, &model.Audit{
			Action:     "send message",
			CallParams: fmt.Sprintf("msg %+v", msg),
		}); txErr != nil {
			return txErr
		}

		return txErr
	})
}
