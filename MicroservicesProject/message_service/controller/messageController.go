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

func (mc *MessageController) GetAllById(ctx context.Context, request *pb.GetAllByIdRequest) (*pb.GetAllByIdResponse, error) {
	fmt.Println(request.Id)
	messages, err := mc.service.GetAllById(request.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllByIdResponse{
		Messages: []*pb.Message{},
	}
	for _, message := range messages {
		current := mapMessage(message)
		response.Messages = append(response.Messages, current)
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
	fmt.Println(id)
	status := request.Status
	fmt.Println(status)
	_id, err := mc.service.ChangeMessageStatus(status, id)
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
