package main

import (
	"rest-api-golang-jwt/controllers"
	"rest-api-golang-jwt/database"
	"rest-api-golang-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	// Init Database
	database.Connect("root:@tcp(localhost:3306)/golang_jwt?parseTime=true")
	database.Migrate()

	router := initRoter()
	router.Run(":8080")
}

func initRoter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)

		secure := api.Group("/secured").Use(middlewares.Auth())
		{
			secure.GET("/ping", controllers.Ping)
		}

	}
	return router
}
