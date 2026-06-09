package main

import (
	"thaibev-assignment/backend/database"
	"thaibev-assignment/backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.GET("/questions", handlers.GetQuestions)
		api.POST("/questions", handlers.CreateQuestion)
		api.DELETE("/questions/:id", handlers.DeleteQuestion)
	}

	r.Run(":8080")
}
