package main

import (
	"Canteen/configuration"
	"Canteen/helper"
	"Canteen/middleware"
	"Canteen/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	configuration.LoadConfig("./config.ini")
	helper.MysqlConnect()
}

func main() {
	app := gin.Default()
	app.Use(middleware.Cors())

	router.LoginRegister(app)
	router.SuggestRegister(app)

	err := app.Run("0.0.0.0:" + configuration.Configs.ServerPort)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
