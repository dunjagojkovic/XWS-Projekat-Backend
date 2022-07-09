package controller

import (
	pb "common/proto/message_service"
	"context"
	"fmt"
	"messageS/model"
	"messageS/service"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageController struct {
	pb.UnimplementedMessageServiceServer
	service      *service.MessageService
	CustomLogger *CustomLogger
}

func NewMessageController(service *service.MessageService) *MessageController {
	CustomLogger := NewCustomLogger()
	return &MessageController{
		service:      service,
		CustomLogger: CustomLogger,
	}

}

func (mc *MessageController) GetMessages(ctx context.Context, request *pb.GetChatsRequest) (*pb.GetMessagesResponse, error) {
	fmt.Println(request.Id)
	messages, err := mc.service.GetMessages(request.Id)
	if err != nil {
		mc.CustomLogger.ErrorLogger.Error("Get all messages for user unsuccessful")
		return nil, err
	}
	response := &pb.GetMessagesResponse{
		Messages: []*pb.Message{},
	}
	for _, message := range messages {
		current := mapMessage(&message)
		response.Messages = append(response.Messages, current)
	}
	mc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(messages)) + " messages")

	return response, nil
}

func (mc *MessageController) GetChats(ctx context.Context, request *pb.GetChatsRequest) (*pb.GetChatsResponse, error) {
	fmt.Println(request.Id)
	chats, list, err := mc.service.GetChats(request.Id)
	fmt.Println(list)
	if err != nil {
		return nil, err
	}
	response := &pb.GetChatsResponse{
		Chats: []*pb.Chat{},
		List:  list,
	}
	for _, chat := range chats {
		current := mapChat(chat)
		response.Chats = append(response.Chats, current)
	}
	return response, nil
}

func (mc *MessageController) CreateMessage(ctx context.Context, request *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {

	message := mapNewMessage(request.CreatedMessage)
	id, _, err := mc.service.CreateMessage(message)
	if err != nil {
		mc.CustomLogger.ErrorLogger.Error("ObjectId not created with ID:" + message.Id.Hex())
		return nil, err
	}
	response := &pb.CreateMessageResponse{
		Id: id.Hex(),
	}

	mc.CustomLogger.SuccessLogger.Info("Message with ID: " + message.Id.Hex() + " created succesfully by user with ID: " + message.Sender.Hex() + " and message status: " + message.Status)
	return response, nil

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
