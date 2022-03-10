package main

import (
	"fmt"
	//"log"
	"todo/app/controllers"
	"todo/app/models"
	//"todo/config"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}