package model

import "time"

type Chat struct {
	Name  string
	Users []User
}

type User struct {
	Id   int64
	Name string
}

type Message struct {
	User      User
	Text      string
	CreatedAt time.Time
}
