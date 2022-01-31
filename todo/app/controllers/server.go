package controllers

import (
	"net/http"
	"todo/config"
)

func StartMainServer() error {
	http.HandleFunc("/",top)
	return http.ListenAndServe(":"+config.Config.Port,nil)
}