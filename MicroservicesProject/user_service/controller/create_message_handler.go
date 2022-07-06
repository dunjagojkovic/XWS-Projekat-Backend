package controller

import (
	events "common/saga/create_message"
	saga "common/saga/messaging"
	"userS/service"
)

type CreateMessageCommandHandler struct {
	userService       *service.UserService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateMessageCommandHandler(userService *service.UserService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateMessageCommandHandler, error) {
	o := &CreateMessageCommandHandler{
		userService:       userService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateMessageCommandHandler) handle(command *events.CreateMessageCommand) {
	reply := events.CreateMessageReply{Message: command.Message}

	switch command.Type {
	case events.CheckBlocking:
		isBlocked := handler.userService.CheckBlocking(command.Message.Receiver, command.Message.Sender)
		if isBlocked == true {
			reply.Type = events.UserBlocked
			break
		}
		reply.Type = events.UserNotBlocked

	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
