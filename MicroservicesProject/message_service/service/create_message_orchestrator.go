package service

import (
	events "common/saga/create_message"
	saga "common/saga/messaging"
	"messageS/model"
)

type CreateMessageOrchestrator struct {
	commandPublisher saga.Publisher
	replySubscriber  saga.Subscriber
}

func NewCreateMessageOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) (*CreateMessageOrchestrator, error) {
	o := &CreateMessageOrchestrator{
		commandPublisher: publisher,
		replySubscriber:  subscriber,
	}
	err := o.replySubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *CreateMessageOrchestrator) Start(message *model.Message, idChat string) error {
	event := &events.CreateMessageCommand{
		Type: events.CheckBlocking,
		Message: events.MessageDetails{
			Id:       message.Id.Hex(),
			Sender:   message.Sender.Hex(),
			Receiver: message.Receiver.Hex(),
			ChatId:   idChat,
		},
	}

	return o.commandPublisher.Publish(event)
}

func (o *CreateMessageOrchestrator) handle(reply *events.CreateMessageReply) {
	command := events.CreateMessageCommand{Message: reply.Message}
	command.Type = o.nextCommandType(reply.Type)
	if command.Type != events.UnknownCommand {
		_ = o.commandPublisher.Publish(command)
	}
}

func (o *CreateMessageOrchestrator) nextCommandType(reply events.CreateMessageReplyType) events.CreateMessageCommandType {
	switch reply {
	case events.UserBlocked:
		return events.CancelMessage
	case events.UserNotBlocked:
		return events.ApproveMessage
	default:
		return events.UnknownCommand
	}
}
