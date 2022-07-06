package controller

import (
	events "common/saga/create_message"
	saga "common/saga/messaging"
	"messageS/service"
)

type CreateMessageCommandHandler struct {
	messageService    *service.MessageService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateMessageCommandHandler(messageService *service.MessageService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateMessageCommandHandler, error) {
	o := &CreateMessageCommandHandler{
		messageService:    messageService,
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
	case events.ApproveMessage:
		_, err := handler.messageService.ChangeMessageStatus("Sent", command.Message.Id, command.Message.ChatId)
		if err != nil {
			return
		}
		reply.Type = events.MessageApproved
	case events.CancelMessage:
		_, err := handler.messageService.ChangeMessageStatus("Canceled", command.Message.Id, command.Message.ChatId)
		if err != nil {
			return
		}
		reply.Type = events.MessageCancelled
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
