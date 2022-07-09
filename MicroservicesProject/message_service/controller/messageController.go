package controller

import (
	pb "common/proto/message_service"
	"common/tracer"
	"context"
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
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetMessages")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	messages, err := mc.service.GetMessages(ctx, request.Id)
	if err != nil {
		mc.CustomLogger.ErrorLogger.Error("Get all messages for user unsuccessful")
		return nil, err
	}
	response := &pb.GetMessagesResponse{
		Messages: []*pb.Message{},
	}
	for _, message := range messages {
		current := mapMessage(ctx, &message)
		response.Messages = append(response.Messages, current)
	}
	mc.CustomLogger.SuccessLogger.Info("Found " + strconv.Itoa(len(messages)) + " messages")

	return response, nil
}

func (mc *MessageController) GetChats(ctx context.Context, request *pb.GetChatsRequest) (*pb.GetChatsResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER GetChats")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	chats, list, err := mc.service.GetChats(ctx, request.Id)

	if err != nil {
		return nil, err
	}
	response := &pb.GetChatsResponse{
		Chats: []*pb.Chat{},
		List:  list,
	}
	for _, chat := range chats {
		current := mapChat(ctx, chat)
		response.Chats = append(response.Chats, current)
	}
	return response, nil
}

func (mc *MessageController) CreateMessage(ctx context.Context, request *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER CreateMessage")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	message := mapNewMessage(ctx, request.CreatedMessage)
	id, _, err := mc.service.CreateMessage(ctx, message)
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
	span := tracer.StartSpanFromContext(ctx, "CONTROLLER ChangeMessageStatus")
	defer span.Finish()

	id := request.Id
	chatId := request.IdChat
	status := request.Status

	ctx = tracer.ContextWithSpan(context.Background(), span)
	_id, err := mc.service.ChangeMessageStatus(ctx, status, id, chatId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateMessageResponse{
		Id: _id.Hex(),
	}, nil

}

func mapNewMessage(ctx context.Context, messagePb *pb.CreateMessage) *model.Message {
	span := tracer.StartSpanFromContext(ctx, "mapNewMessage")
	defer span.Finish()

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

func mapMessage(ctx context.Context, message *model.Message) *pb.Message {
	span := tracer.StartSpanFromContext(ctx, "mapMessage")
	defer span.Finish()

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

func mapChat(ctx context.Context, chat *model.Chat) *pb.Chat {
	span := tracer.StartSpanFromContext(ctx, "mapChat")
	defer span.Finish()

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
