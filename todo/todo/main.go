package main

import (
	"fmt"
	"todo/app/controllers"
	"todo/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}