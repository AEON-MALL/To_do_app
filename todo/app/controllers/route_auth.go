package controllers

import (
	"log"
	"net/http"
	"todo/app/models"
)

func signup(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
<<<<<<< HEAD
		_, err := session(w, r)
		if err != nil{
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		}else{
			http.Redirect(w, r, "/todos", http.StatusFound)
=======
		_ , err := session(w,r)
		if err != nil{
			generateHTML(w,nil,"layout","public_navbar","signup")
		}else{
			http.Redirect(w,r,"/todos",http.StatusFound)
>>>>>>> origin/main
		}
	}else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil{
			log.Println(err)
		}
		user := models.User{
			Name: r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err !=nil {
			log.Println(err)
		}

<<<<<<< HEAD
		http.Redirect(w, r, "/", http.StatusFound)
=======
		http.Redirect(w,r,"/",http.StatusFound)
>>>>>>> origin/main
	}
}

func login(w http.ResponseWriter, r *http.Request){
<<<<<<< HEAD
	_, err := session(w,r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	}else{
		http.Redirect(w, r, "/todos", http.StatusFound)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request){
=======
	_ , err := session(w,r)
	if err != nil {
		generateHTML(w,nil,"layout","public_navbar","login")
	}else{
		http.Redirect(w,r,"/todos",http.StatusFound)
	}
}

func auhtenticate(w http.ResponseWriter, r *http.Request){
>>>>>>> origin/main
	err := r.ParseForm()
	if err != nil{
		log.Println(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
<<<<<<< HEAD
		http.Redirect(w, r, "/login", http.StatusFound)
=======
		http.Redirect(w,r,"/login",http.StatusFound)
>>>>>>> origin/main
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")){
		session, err := user.CreateSession()
		if err != nil{
			log.Println(err)
		}

		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

<<<<<<< HEAD
		http.Redirect(w, r, "/", http.StatusFound)
	}else{
		http.Redirect(w, r, "/login", http.StatusFound)
=======
		http.Redirect(w,r,"/",http.StatusFound)
	}else{
		http.Redirect(w,r,"/login",http.StatusFound)
>>>>>>> origin/main
	}
}

func logout(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie{
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
<<<<<<< HEAD
	http.Redirect(w, r, "/login", http.StatusFound)
=======
	http.Redirect(w,r,"/login",http.StatusFound)
>>>>>>> origin/main
}