package main

import (
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/abubakar-sadiq001/docs"
	"github.com/abubakar-sadiq001/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API Documentation
// @version 1
// @Description sample description

// @contact.name scripter_js
// @contact.url githib.com/abubakar-sadiq001
// @contact.email scripter@demo.com

// @securityDefinitions.apikey bearertoken
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	user := v1.Group("/users")
	{
		user.GET("/", handlers.GetUsers)
		user.POST("/", handlers.CreateUser)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
