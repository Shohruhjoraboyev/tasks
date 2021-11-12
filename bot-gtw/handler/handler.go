package handler

import (
	"bot-gtw/proto/sender"
	"context"

	"github.com/gin-gonic/gin"
)

type message struct {
	Text     string `json:"text"`
	Priority string `json:"priority"`
}

var cl sender.SendClient

var newMsg message

func Init(c sender.SendClient) {
	r := gin.Default()
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//send := r.Group("/send")
	r.POST("/", SendMessage)
	r.Run(":8090")
	cl = c
	//SendMessage(&gin.Context{})
}
func SendMessage(c *gin.Context) {
	if err := c.BindJSON(&newMsg); err != nil {
		return
	}
	cl.SendMsg(context.Background(), &sender.SendReq{Text: newMsg.Text, Priority: newMsg.Priority})
}
