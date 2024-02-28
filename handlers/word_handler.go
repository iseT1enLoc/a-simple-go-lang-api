package handlers

import (
	"errors"
	"net/http"
	"simpleAPI/database"
	"simpleAPI/models"

	"github.com/gin-gonic/gin"
)

func GetAllWords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.WordList)
}

func WordById(c *gin.Context) {
	Id := c.Param("id")
	word, err := GetWordById(Id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Can not fing your expected word"})
		return
	}
	c.IndentedJSON(http.StatusOK, word)
}
func GetWordById(Id string) (*models.Word, error) {
	for i, word := range database.WordList {
		if word.WordID == Id {
			return &database.WordList[i], nil

		}
	}
	return nil, errors.New("word not found in database")
}
