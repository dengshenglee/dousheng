package service

import (
	"fmt"
	"log"
	"mini-douyin/config"
	"testing"
	"time"
)

func TestMessageServiceImpl_SendMessage(t *testing.T) {
	err := GetMessageServiceInstance().SendMessage(1, 7, "1 号用户发送消息给 7 号用户", 1)
	if err == nil {
		log.Println("SendMessage Service 正常")
	}
}

func TestMessageServiceImpl_MessageChat(t *testing.T) {
	oldTime, err := time.Parse(config.GO_STARTER_TIME, "2023-03-20 15:04:05")
	if err != nil {
		log.Println(err)
	}
	chat, _ := GetMessageServiceInstance().MessageChat(1, 7, oldTime)
	for _, msg := range chat {
		log.Println(fmt.Sprintf("%+v", msg))
	}
}

func TestMessageServiceImpl_LatestMessage(t *testing.T) {
	message, _ := GetMessageServiceInstance().LatestMessage(1, 7)
	log.Println(fmt.Sprintf("%+v", message))
}
