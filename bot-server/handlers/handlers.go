package handlers

import (
	"bot-gtwy/model"
	"bot-gtwy/proto/sender"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	sender.UnimplementedSendServer
}
type message struct {
	Text     string `json:"text"`
	Priority string `json:"priority"`
}

var chat_id = -1001724916144

func (s *Server) SendMsg(ctx context.Context, req *sender.SendReq) (*sender.SendResp, error) {
	now = time.Now().Unix()
	timeout = time.Now().Unix() - lastExec
	newMsg = message{
		Text:     req.GetText(),
		Priority: req.GetPriority(),
	}
	msg := SortByPriority(newMsg)
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
				Send(chat_id, mesge.Text)
				prio[i] = message{}
				log.Println(msg)
			}
		}
	}
	return &sender.SendResp{Status: "message send"}, nil
}

const botUrl = "https://api.telegram.org/bot2066190858:AAH963KboIt1760x9RBHjESSzDAbhB_Rcu8"

func Send(chat_id int, text string) {
	botMessage := model.BotMessage{ChatId: chat_id, Text: text}
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

func SortByPriority(m message) [][]message {

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

var lastSent, now, timeout, lastExec, nowSend, timeoutSend int64
var newMsg message

func SendMessage() {
	c := gin.Context{}
	now = time.Now().Unix()
	timeout = time.Now().Unix() - lastExec
	if err := c.BindJSON(&newMsg); err != nil {
		return
	}
	msg := SortByPriority(newMsg)
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
				Send(chat_id, mesge.Text)
				prio[i] = message{}
				log.Println(msg)

				c.IndentedJSON(http.StatusOK, "'"+mesge.Text+"'"+" message was sent")
			}
		}
	}
}
