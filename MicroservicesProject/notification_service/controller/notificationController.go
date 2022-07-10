package controller

import (
	"context"
	"notificationS/model"
	"notificationS/service"

	pb "common/proto/notification_service"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationController struct {
	pb.UnimplementedNotificationServiceServer
	service *service.NotificationService
}

func NewNotificationController(service *service.NotificationService) *NotificationController {
	return &NotificationController{
		service: service,
	}

}

func (pc *NotificationController) GetByUserId(ctx context.Context, request *pb.GetByUserIdRequest) (*pb.GetByUserIdResponse, error) {
	objID, err := primitive.ObjectIDFromHex(request.Id)

	notifications, err := pc.service.GetByUserId(objID)
	if err != nil {
		return nil, err
	}
	response := &pb.GetByUserIdResponse{
		Notifications: []*pb.Notification{},
	}
	for _, notification := range notifications {
		current := mapNotification(&notification)
		response.Notifications = append(response.Notifications, current)
	}
	return response, nil
}

func mapNotification(notification *model.Notification) *pb.Notification {
	notificationPb := &pb.Notification{
		Id:     notification.Id.Hex(),
		Text:   notification.Text,
		Time:   timestamppb.New(notification.Time),
		UserId: notification.UserId.Hex(),
		Read:   notification.Read,
	}

	return notificationPb
}

func mapNewNotification(notificationPb *pb.CreateNotificationRequest) *model.Notification {

	objID, _ := primitive.ObjectIDFromHex(notificationPb.UserId)
	notification := &model.Notification{
		Id:     primitive.NewObjectID(),
		Text:   notificationPb.Text,
		Time:   notificationPb.Time.AsTime(),
		UserId: objID,
		Read:   notificationPb.Read,
	}

	return notification
}

func (pc *NotificationController) ChangeStatus(ctx context.Context, request *pb.ChangeStatusRequest) (*pb.GetByUserIdRequest, error) {

	id := request.Id
	status := request.Status
	objectId, err := primitive.ObjectIDFromHex(id)
	_id, err := pc.service.ChangeStatus(objectId, status)

	if err != nil {
		return nil, err
	}
	response := &pb.GetByUserIdRequest{
		Id: _id.Hex(),
	}

	return response, nil

}

func (pc *NotificationController) CreateNotification(ctx context.Context, request *pb.CreateNotificationRequest) (*pb.GetByUserIdRequest, error) {

	notification := mapNewNotification(request)
	id, err := pc.service.CreateNotification(notification)
	if err != nil {
		return nil, err
	}
	return &pb.GetByUserIdRequest{
		Id: id.Hex(),
	}, nil

}
