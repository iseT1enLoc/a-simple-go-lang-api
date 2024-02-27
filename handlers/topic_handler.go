package handlers

import (
	"net/http"

	"simpleAPI/database"
	"simpleAPI/models"

	"github.com/gin-gonic/gin"

	"errors"
)

func GetAllTopics(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.TopicList)
}
func CreateTopic(c *gin.Context) {
	var new_topic models.Topic

	if err := c.BindJSON(&new_topic); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "something bad happened while creating"})
		return
	}
	database.TopicList = append(database.TopicList, new_topic)
	c.IndentedJSON(http.StatusCreated, new_topic)
}

func TopicById(c *gin.Context) {
	id := c.Param("id")
	topic, err := GetTopicById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Page not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, topic)
}
func GetTopicById(Id string) (*models.Topic, error) {
	for i, topic := range database.TopicList {
		if topic.Id == Id {
			return &database.TopicList[i], nil
		}
	}
	return nil, errors.New("topic not found")
}
