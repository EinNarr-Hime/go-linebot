package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

var client *linebot.Client

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file load faild.")
	}
	accessId := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACESS_SECRET")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set.")
	}
	bot, err := linebot.New(
		accessId,
		accessSecret,
	)
	if err != nil {
		log.Fatal(err)
	}
	client = bot
	router := gin.New()
	router.Use(gin.Logger())

	// LINE Messaging API ルーティング
	router.POST("/callback", replyMessage)
	router.Run(":" + port)
}

func replyMessage(c *gin.Context) {
	events, err := client.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			log.Print(err)
		}
		return
	}
	for _, event := range events {
		// イベントがメッセージの受信だった場合
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				// 返事の内容
				replyMessage := message.Text
				// そのまま返事
				_, err = client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}