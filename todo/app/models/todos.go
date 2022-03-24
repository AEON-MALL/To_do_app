package models

import (
	"log"
	"time"
)

type Todo struct {
	ID int
	Content string
	UserID int
	CreatedAt time.Time
}

<<<<<<< HEAD
func (u *User) CreateTodo(content string) (err error){
	cmd := `insert into todos (
		content, 
		user_id, 
		created_at) values (?, ?, ?)`
=======
func (u *User) CreateTodo(content string)(err error){
	cmd := `insert into todos (content, user_id, created_at) values (?,?,?)`
>>>>>>> origin/main

	_, err = Db.Exec(cmd, content,u.ID , time.Now())
	if err != nil{
		log.Fatalln(err)
	}
	return err
}

func GetTodo(id int)(todo Todo, err error){
<<<<<<< HEAD
	cmd := `select id, content, user_id, created_at 
	from todos where id = ?`
=======
	cmd := `select id, content, user_id,created_at from todos where id = ?`
>>>>>>> origin/main
	todo = Todo{}

	err = Db.QueryRow(cmd,id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo,err
}

func GetTodos() (todos []Todo, err error){
<<<<<<< HEAD
	cmd := `select id, content, user_id, created_at from todos`
=======
	cmd := `select id,content,user_id,created_at from todos`
>>>>>>> origin/main
	rows, err := Db.Query(cmd)
	if err != nil{
		log.Fatalln(err)
	}
	for rows.Next(){
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)

		if err != nil{
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error){
<<<<<<< HEAD
	cmd := `select id, content, user_id, created_at 
	from todos where user_id = ?`
=======
	cmd := `select id, content, user_id, created_at from todos where user_id = ?`
>>>>>>> origin/main

	rows, err := Db.Query(cmd, u.ID)
	if err != nil{
		log.Fatalln(err)
	}
	for rows.Next(){
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil{
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func(t *Todo) UpdateTodo() error {
<<<<<<< HEAD
	cmd := `update todos set content = ?, user_id = ? where id = ?`
=======
	cmd := `update todos set content = ? , user_id = ? where id = ?`
>>>>>>> origin/main
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil{
		log.Fatalln(err)
	}
	return err
}

<<<<<<< HEAD
func(t *Todo) DeleteTodo() error{
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
=======
func(t *Todo)DeleteTodo() error{
	cmd := `delete from todos where id = ?`
	_ , err = Db.Exec(cmd, t.ID)
>>>>>>> origin/main
	if err != nil {
		log.Fatalln(err)
	}
	return err
}