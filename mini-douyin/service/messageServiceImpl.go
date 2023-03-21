package service

import (
	"errors"
	"fmt"
	"log"
	"mini-douyin/config"
	"mini-douyin/dao"
	"sync"
	"time"
)

type MessageServiceImpl struct {
}

var (
	messageServiceImp  *MessageServiceImpl
	messageServiceOnce sync.Once
)

func GetMessageServiceInstance() *MessageServiceImpl {
	messageServiceOnce.Do(func() {
		messageServiceImp = &MessageServiceImpl{}
	})
	return messageServiceImp
}

var _ MessageService = (*MessageServiceImpl)(nil)

// LatestMessage implements MessageService
func (*MessageServiceImpl) LatestMessage(loginUserId int64, targetUserId int64) (LatestMessage, error) {
	plainMessage, err := dao.LatestMessage(loginUserId, targetUserId)
	if err != nil {
		log.Println("LatestMessage Service:", err)
		return LatestMessage{}, err
	}
	var latestMessage LatestMessage
	latestMessage.message = plainMessage.MsgContent
	if plainMessage.UserId == loginUserId {
		//current user send the latest message
		latestMessage.msgType = 1
	} else {
		//current user receive the latest message
		latestMessage.msgType = 0
	}
	return latestMessage, nil
}

// MessageChat implements MessageService
func (*MessageServiceImpl) MessageChat(loginUserId int64, targetUserId int64, latestTime time.Time) ([]Message, error) {
	message := make([]Message, 0, config.VIDEO_INIT_NUM_PER_AUTHOR)
	plainMessages, err := dao.MessageChat(loginUserId, targetUserId, latestTime)
	if err != nil {
		log.Println("MessageChat Service:", err)
		return nil, err
	}
	err = getRespMessage(&message, &plainMessages)
	if err != nil {
		log.Println("getRespMessage:", err)
		return nil, err
	}
	return message, nil
}

// SendMessage implements MessageService
func (*MessageServiceImpl) SendMessage(fromUserId int64, toUserId int64, content string, actionType int64) error {
	var err error
	switch actionType {
	//actionType= 1 SendMessage
	case 1:
		err = dao.SendMessage(fromUserId, toUserId, content, actionType)
	default:
		log.Println(fmt.Sprintf("未定义 actionType=%d", actionType))
		return errors.New(fmt.Sprintf("未定义 actionType=%d", actionType))
	}
	return err
}

// 返回 message list 接口所需的 Message 结构体
func getRespMessage(messages *[]Message, plainMessages *[]dao.Message) error {
	for _, tmpMessage := range *plainMessages {
		var message Message
		combineMessage(&message, &tmpMessage)
		*messages = append(*messages, message)
	}
	return nil
}

func combineMessage(message *Message, plainMessage *dao.Message) error {
	message.Id = plainMessage.Id
	message.UserId = plainMessage.UserId
	message.ReceiverId = plainMessage.ReceiverId
	message.MsgContent = plainMessage.MsgContent
	message.CreatedAt = plainMessage.CreatedAt.Unix()
	return nil
}
