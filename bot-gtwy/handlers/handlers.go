package handlers

import (
	"bot-gtwy/model"
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const botUrl = "https://api.telegram.org/bot2066190858:AAH963KboIt1760x9RBHjESSzDAbhB_Rcu8"

type Message struct {
	Text     string `json:"text"`
	Priority string `json:"priority"`
}

var (
	message1 = Message{"hello", "low"}
	message2 = Message{"salom", "low"}
	message3 = Message{"hi", "medium"}
	message4 = Message{"how r u", "high"}
	message5 = Message{"bonjour", "high"}
)
var messages = []Message{
	message1,
	message2,
	message3,
	message4,
	message5,
}

func SendMessage() error {

	chat_id := -1001724916144
	msg := SortByPriority()
	for _, prio := range msg {
		for _, message := range prio {
			botMessage := model.BotMessage{ChatId: chat_id, Text: message.Text}
			buf, err := json.Marshal(&botMessage)
			if err != nil {
				return err
			}
			_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
			if err != nil {
				return err
			}
			time.Sleep(5 * time.Second)
		}
	}
	return nil
}

func SortByPriority() [][]Message {

	medium := []Message{}
	low := []Message{}
	high := []Message{}
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

	msg := [][]Message{
		high,
		medium,
		low,
	}
	return msg
}
