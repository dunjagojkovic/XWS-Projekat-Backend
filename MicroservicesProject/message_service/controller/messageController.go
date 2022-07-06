package controller

import (
	pb "common/proto/message_service"
	"context"
	"fmt"
	"messageS/model"
	"messageS/service"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageController struct {
	pb.UnimplementedMessageServiceServer
	service *service.MessageService
}

func NewMessageController(service *service.MessageService) *MessageController {
	return &MessageController{
		service: service,
	}

}

func (mc *MessageController) GetMessages(ctx context.Context, request *pb.GetChatsRequest) (*pb.GetMessagesResponse, error) {
	fmt.Println(request.Id)
	messages, err := mc.service.GetMessages(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetMessagesResponse{
		Messages: []*pb.Message{},
	}
	for _, message := range messages {
		current := mapMessage(&message)
		response.Messages = append(response.Messages, current)
	}
	return response, nil
}

func (mc *MessageController) GetChats(ctx context.Context, request *pb.GetChatsRequest) (*pb.GetChatsResponse, error) {
	fmt.Println(request.Id)
	chats, err := mc.service.GetChats(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetChatsResponse{
		Chats: []*pb.Chat{},
	}
	for _, chat := range chats {
		current := mapChat(chat)
		response.Chats = append(response.Chats, current)
	}
	return response, nil
}

func (mc *MessageController) CreateMessage(ctx context.Context, request *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {

	message := mapNewMessage(request.CreatedMessage)
	id, err := mc.service.CreateMessage(message)
	if err != nil {
		return nil, err
	}
	return &pb.CreateMessageResponse{
		Id: id.Hex(),
	}, nil

}

func (mc *MessageController) ChangeMessageStatus(ctx context.Context, request *pb.ChangeMessageStatusRequest) (*pb.CreateMessageResponse, error) {

	id := request.Id
	chatId := request.IdChat
	fmt.Println(id)
	status := request.Status
	fmt.Println(status)
	_id, err := mc.service.ChangeMessageStatus(status, id, chatId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateMessageResponse{
		Id: _id.Hex(),
	}, nil

}

func mapNewMessage(messagePb *pb.CreateMessage) *model.Message {

	receiverId, _ := primitive.ObjectIDFromHex(messagePb.Receiver)
	senderId, _ := primitive.ObjectIDFromHex(messagePb.Sender)
	//layout := "2006-01-02 15:04:05"
	//dt, _ := time.Parse(layout, messagePb.Time)
	message := &model.Message{
		Id:       primitive.NewObjectID(),
		Text:     messagePb.Text,
		Sender:   senderId,
		Receiver: receiverId,
		Time:     messagePb.Time.AsTime(),
		Status:   messagePb.Status,
	}

	return message
}

func mapMessage(message *model.Message) *pb.Message {
	messagePb := &pb.Message{
		Id:       message.Id.Hex(),
		Text:     message.Text,
		Sender:   message.Sender.Hex(),
		Receiver: message.Receiver.Hex(),
		Time:     timestamppb.New(message.Time),
		Status:   message.Status,
	}

	return messagePb
}

func mapChat(chat *model.Chat) *pb.Chat {
	chatPb := &pb.Chat{
		Id:         chat.Id.Hex(),
		FirstUser:  chat.FirstUser.Hex(),
		SecondUser: chat.SecondUser.Hex(),
		Messages:   make([]*pb.Message, 0),
	}

	for _, message := range chat.Messages {

		messagePb := *&pb.Message{
			Id:       message.Id.Hex(),
			Text:     message.Text,
			Sender:   message.Sender.Hex(),
			Receiver: message.Receiver.Hex(),
			Time:     timestamppb.New(message.Time),
			Status:   message.Status,
		}

		chatPb.Messages = append(chatPb.Messages, &messagePb)
	}

	return chatPb
}
