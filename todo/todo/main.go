package main

import (
	"fmt"
	"log"
	//"todo/app/controllers"
	"todo/app/models"
	//"todo/config"
)

func main() {
	fmt.Println(models.Db)

	//controllers.StartMainServer()

	user, _  := models.GetUserByEmail("test@example.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil{
		log.Println(err)
	}
	fmt.Println(session)
}