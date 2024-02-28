package handlers

import (
	"errors"
	"net/http"
	"simpleAPI/database"
	"simpleAPI/models"

	"github.com/gin-gonic/gin"
)

func GetAllWords(c *gin.Context) {
	Id := c.Param("topic_id")
	if len(Id) == 0 {
		c.IndentedJSON(http.StatusOK, database.WordList)
		return
	}
	//Id, _ := strconv.Atoi(temp)
	var wordlist []models.Word
	for _, word := range database.WordList {
		if word.TopicID == Id {
			wordlist = append(wordlist, word)
		}
	}
	if len(wordlist) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "do not have that words"})
		return
	}
	c.IndentedJSON(http.StatusOK, wordlist)
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
func CreateWord(c *gin.Context) {
	var word models.Word

	if err := c.BindJSON(&word); err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Can not create word"})
		return
	}
	database.WordList = append(database.WordList, word)
	c.IndentedJSON(http.StatusOK, word)
}
