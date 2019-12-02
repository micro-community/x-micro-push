package broker

import (
	"encoding/json"
	"fmt"
	proto "github.com/micro-community/x-micro-push/proto"
	"github.com/micro-community/x-micro-push/session"
	"github.com/micro/go-micro/util/log"
)

//Message  Struct
type Message struct {
	Data EventMessage
}

//EventMessage ...
type EventMessage struct {
	Topic   string
	Message *proto.MessageEntity
}

var (
	messageBroadcast = make(chan EventMessage, 512)
)

//GetMessageQueue to push
func GetMessageQueue() chan EventMessage {
	return messageBroadcast
}

// PushMessage 推送消息到前端
func PushMessage() {
	log.Info(nil, "start push message")

	for {
		msg := <-messageBroadcast
		body := []byte(msg.Message.GetBody())
		var message interface{}
		json.Unmarshal(body, &message)

		data := map[string]interface{}{
			"Topic":   msg.Topic,
			"Message": message,
		}
		fmt.Println("message:", data)
		session.SendMessage(data, msg.Message.GetUsers())
	}
}

//Push ...
func Push(msg EventMessage) {
	messageBroadcast <- msg
}
