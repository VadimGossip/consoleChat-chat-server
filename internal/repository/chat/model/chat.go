package model

import (
	"time"
)

type Message struct {
	From      string
	Text      string
	CreatedAt time.Time
}

type Chat struct {
	ID        int64
	Users     []string
	CreatedAt time.Time
	Messages  []Message
}
