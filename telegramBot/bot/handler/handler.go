package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	_ "tasks/telegramBot/bot/docs"
	"tasks/telegramBot/bot/models"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const botUrl = "https://api.telegram.org/bot2066190858:AAH963KboIt1760x9RBHjESSzDAbhB_Rcu8"

type message struct {
	Text     string `json:"text"`
	Priority string `json:"priority"`
}

func Init() {

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	send := r.Group("/send")
	send.POST("", SendMessage)
	r.Run()
}

var arr []message = nil

//@Summary SendMessages
//@Description Send all messages
//@ID send-message
//@Accept json
//@Produce json
//@Param message body message true "new message"
//@Success 200 {string} string ""
//@Failure 404 {object} models.HTTPError
//@Router /send [post]
var lastSent, now, timeout, lastExec, nowSend, timeoutSend int64
var newMsg message

func SendMessage(c *gin.Context) {
	now = time.Now().Unix()
	timeout = time.Now().Unix() - lastExec
	if err := c.BindJSON(&newMsg); err != nil {
		return
	}
	msg := sortByPriority(newMsg)
	log.Println(msg)

	if timeout < 5 {
		time.Sleep((5 - time.Duration(timeout)) * time.Second)
	}
	lastExec = now
	chat_id := -1001724916144

	for _, prio := range msg {
		for i, mesge := range prio {
			if len(mesge.Text) > 0 {
				nowSend = time.Now().Unix()
				timeoutSend = time.Now().Unix() - lastSent
				log.Println(timeoutSend)
				if timeoutSend < 5 {
					time.Sleep((5 - time.Duration(timeout)) * time.Second)
				}
				lastSent = nowSend
				send(chat_id, mesge.Text)
				prio[i] = message{}
				log.Println(msg)

				c.IndentedJSON(http.StatusOK, "'"+mesge.Text+"'"+" message was sent")
			}
		}
	}
}
func send(chat_id int, text string) {
	botMessage := models.BotMessage{ChatId: chat_id, Text: text}
	buf, err := json.Marshal(&botMessage)
	if err != nil {
		log.Println(err)
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Fatal(err)
	}
}

var (
	medium = []message{}
	low    = []message{}
	high   = []message{}
)

func sortByPriority(m message) [][]message {

	switch m.Priority {
	case "high":
		high = append(high, m)
	case "medium":
		medium = append(medium, m)
	case "low":
		low = append(low, m)
	}
	msg := [][]message{
		high,
		medium,
		low,
	}
	return msg
}
