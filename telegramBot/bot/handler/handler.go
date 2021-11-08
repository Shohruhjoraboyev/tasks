package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"tasks/telegramBot/bot/models"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "tasks/telegramBot/bot/docs"
)

const botUrl = "https://api.telegram.org/bot2066190858:AAH963KboIt1760x9RBHjESSzDAbhB_Rcu8"

type message struct {
	Text     string `json:"text"`
	Priority string `json:"priority"`
}

var (
	message1 = message{"hello", "low"}
	message2 = message{"salom", "low"}
	message3 = message{"hi", "medium"}
	message4 = message{"how r u", "high"}
	message5 = message{"bonjour", "high"}
)
var messages = []message{
	message1,
	message2,
	message3,
	message4,
	message5,
}

func Init() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	send := r.Group("/send")
	send.POST("/sendMessage", SendMessage)
	get := r.Group("/get")
	get.GET("/", Messageall)

	r.Run()

}

//@Summary SendMessages
//@Description Send all messages
//@ID send-message
//@Accept json
//Produce json
//@Success 200 {string} string ""
//@Failure 404 {object} models.HTTPError
//@Router /sendMessage [post]
func SendMessage(c *gin.Context) {

	cl := &http.Client{}

	chat_id := -1001724916144
	msg := sortByPriority()
	for _, prio := range msg {
		for _, message := range prio {
			botMessage := models.BotMessage{ChatId: chat_id, Text: message.Text}
			buf, err := json.Marshal(&botMessage)
			if err != nil {
				log.Println(err)
			}
			_, err = cl.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
			if err != nil {
				log.Println(err)
			}
			time.Sleep(5 * time.Second)
		}
	}
}
