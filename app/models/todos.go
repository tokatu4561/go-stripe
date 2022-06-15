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

func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `select id, content, user_id, create_at
	from todos where id = ?`

	err = DB.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id, create_at
	from todos`

	rows, err := DB.Query(cmd)
	if err != nil {
		log.
		Fatalln(err)
	}

	for rows.Next() {
		var todo Todo 

		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}

	rows.Close()
	
	return todos , err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id, create_at
	from todos where user_id = ?`

	rows, err := DB.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo 

		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}

	rows.Close()

	return todos , err
}