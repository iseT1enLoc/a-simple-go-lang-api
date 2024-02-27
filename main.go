package main

import (
	"simpleAPI/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/topics", handlers.GetAllTopics)
	router.POST("/topics", handlers.CreateTopic)
	router.GET("/topics/:id", handlers.TopicById)
	router.Run("localhost:8080")
}
