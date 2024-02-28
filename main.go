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

	router.GET("/words", handlers.GetAllWords)
	router.GET("/words/:id", handlers.WordById)
	router.POST("/words", handlers.CreateWord)
	router.Run("localhost:8080")
}
