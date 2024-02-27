package handlers

import (
	"net/http"

	"simpleAPI/database"

	"github.com/gin-gonic/gin"
)

func GetAllTopics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.TopicList)
}
