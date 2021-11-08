package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sortByPriority() [][]message {

	medium := []message{}
	low := []message{}
	high := []message{}
	for i, message := range messages {
		switch message.Priority {
		case "high":
			high = append(high, messages[i])
		case "medium":
			medium = append(medium, messages[i])
		case "low":
			low = append(low, messages[i])
		}
	}

	msg := [][]message{
		high,
		medium,
		low,
	}
	return msg
}

//@Summary AllMessages
//@Description Show all messages
//@ID get-all-messages
//@Accept json
//@Produce json
//@Success 200 {array} msg
//@Failure 404 {object} models.HTTPError
//@Router /getMessages [get]

func Messageall(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sortByPriority())
}
