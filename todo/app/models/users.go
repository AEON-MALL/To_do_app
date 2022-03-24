package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos	[]Todo
}

type Session struct {
	ID int
	UUID string
	Email string
	UserID int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error){
<<<<<<< HEAD
	cmd := `insert into users (
		uuid, 
		name, 
		email, 
		password, 
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
=======
	cmd := `insert into users(uuid, name, email, password, created_at) values (?,?,?,?,?)`

	_ , err = Db.Exec(cmd,
>>>>>>> origin/main
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil{
		log.Fatalln(err)
	}
	return err
}

<<<<<<< HEAD
func GetUser(id int) (user User, err error){
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at 
	from users where id = ?`
=======
func GetUser(id int)(user User, err error){
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`
>>>>>>> origin/main
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user,err
}

func (u *User) UpdateUser() (err error){
<<<<<<< HEAD
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name,u.Email,u.ID)
=======
	cmd := `update users set name = ? , email = ? where id = ?`
	_ , err = Db.Exec(cmd, u.Name,u.Email,u.ID)
>>>>>>> origin/main
	if err != nil{
		log.Fatalln(err)
	}
	return err
}

<<<<<<< HEAD
func (u *User) DeleteUser() (err error){
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
=======
func (u *User) DeleteUser()(err error){
	cmd := `delete from users where id = ?`
	_ , err = Db.Exec(cmd, u.ID)
>>>>>>> origin/main
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

<<<<<<< HEAD
func GetUserByEmail(email string) (user User, err error){
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at 
	from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID, 
		&user.Name,
=======
func GetUserByEmail(email string)(user User, err error){
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID, 
>>>>>>> origin/main
		&user.Email, 
		&user.PassWord, 
		&user.CreatedAt)
	
	return user,err
}

func (u *User) CreateSession() (session Session, err error){
	session = Session{}
<<<<<<< HEAD
	cmd1 := `insert into sessions (
		uuid, 
		email, 
		user_id, 
		created_at) values (?, ?, ?, ?)`
	
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
=======
	cmd1 := `insert into sessions (uuid, email, user_id, created_at) values (?, ?, ?, ?)`
	
	_ , err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
>>>>>>> origin/main
	if err != nil {
		log.Println(err)
	}

<<<<<<< HEAD
	cmd2 := `select id, uuid, email, user_id, created_at 
	from sessions where user_id = ? and email = ?`
=======
	cmd2 := `select id, uuid, email, user_id, created_at from sessions where user_id = ? and email = ?`
>>>>>>> origin/main

	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID, 
		&session.UUID, 
		&session.Email, 
		&session.UserID, 
		&session.CreatedAt)
	
	return session, err
}

func (sess *Session) CheckSession() (valid bool, err error){
<<<<<<< HEAD
	cmd := `select id, uuid, email, user_id, created_at 
	from sessions where uuid = ?`
=======
	cmd := `select id, uuid, email, user_id, created_at from sessions where uuid = ?`
>>>>>>> origin/main

	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID, 
		&sess.Email,
		&sess.UserID, 
		&sess.CreatedAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error){
	cmd :=`delete from sessions where uuid = ? `
<<<<<<< HEAD
	_, err = Db.Exec(cmd, sess.UUID)
=======
	_ , err = Db.Exec(cmd, sess.UUID)
>>>>>>> origin/main
	if err != nil{
		log.Println(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error){
	user = User{}
<<<<<<< HEAD
	cmd := `select id, uuid, name, email, created_at 
	from users where id = ?`
=======
	cmd := `select id, uuid, name, email, created_at from users where id = ?`
>>>>>>> origin/main
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	return user, err
}
