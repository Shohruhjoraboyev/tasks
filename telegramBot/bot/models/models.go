package models

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
