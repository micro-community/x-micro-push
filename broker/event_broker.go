package broker

import (
	"context"
	proto "github.com/micro-community/x-push/proto"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/util/log"
)

// EventWorker is an status publisher for the go-micro broker
type EventWorker struct {
	publisher micro.Publisher
}

var (
	eventTopic = proto.EventMessageTopic_EventMessage.String()
)

// NewEventSubPub creates a new broker status publisher
func NewEventSubPub(service micro.Service) *EventWorker {
	return &EventWorker{
		publisher: micro.NewPublisher(eventTopic, service.Client()),
	}
}

//RegisterEventWorker for status
func RegisterEventWorker(service micro.Service) {
	worker := NewEventSubPub(service)
	micro.RegisterSubscriber(eventTopic, service.Server(), worker.EventProcess)

}

//EventProcess Event for MessageBus
func (a *EventWorker) EventProcess(ctx context.Context, eventMessage *proto.MessageEntity) error {
	md, _ := metadata.FromContext(ctx)
	log.Info("Received Event ", md)
	// do something with Event
	eventData := EventMessage{
		Topic:   eventMessage.Topic,
		Message: eventMessage,
	}

	messageBroadcast <- eventData
	return nil
}
