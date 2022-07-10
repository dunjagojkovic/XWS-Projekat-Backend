package main

import (
	notificationGW "common/proto/notification_service"
	"context"
	"fmt"
	"log"
	"net"
	"notificationS/config"
	"notificationS/controller"
	"notificationS/repository"
	"notificationS/service"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.NotificationDBHost, server.config.NotificationDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uri)
	}
	store := repository.NewNotificationStore(client)
	notificationService := service.NewNotificationService(store)

	notificationController := controller.NewNotificationController(notificationService)

	server.startGrpcServer(notificationController)
}

func (server *Server) startGrpcServer(notificationController *controller.NotificationController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	notificationGW.RegisterNotificationServiceServer(grpcServer, notificationController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
