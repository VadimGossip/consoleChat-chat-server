package model

import (
	"time"
)

type User struct {
	Id   int64
	Name string
}

type Message struct {
	User      User
	Text      string
	CreatedAt time.Time
}

type Chat struct {
	ID        int64
	Name      string
	Users     []string
	CreatedAt time.Time
	Messages  []Message
}
