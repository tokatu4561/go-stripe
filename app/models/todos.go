package models

import (
	"time"
	"log"
)

type Todo struct {
	ID int
	Content string
	UserID int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (content, user_id, created_at) values (?, ?, ?)`

	_, err = DB.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}