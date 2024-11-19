package main

import (
	"app/config"
	"app/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := gin.New()
	db := config.DatabaseConn()
	router := app.Group("/api/v1")
	route.Register(router, db)

	app.Run(":5500")

}
