package main

import (
	messageGW "common/proto/message_service"
	"context"
	"fmt"
	"log"
	"messageS/config"
	"messageS/controller"
	"messageS/repository"
	"messageS/service"
	"net"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	config       *config.Config
	CustomLogger *controller.CustomLogger
}

func NewServer(config *config.Config) *Server {
	CustomLogger := controller.NewCustomLogger()
	return &Server{
		config:       config,
		CustomLogger: CustomLogger,
	}
}

func (server *Server) Start() {

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.MessageDBHost, server.config.MessageDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
		server.CustomLogger.ErrorLogger.WithFields(logrus.Fields{"Message db host": server.config.MessageDBHost, "Message db port": server.config.MessageDBPort}).Error("MongoDB initialization for message service failed")
	} else {
		fmt.Println(uri)
		server.CustomLogger.SuccessLogger.Info("MongoDB initialization for message service succesfull")

	}
	store := repository.NewMessageStore(client)
	messageService := service.NewMessageService(store)
	messageController := controller.NewMessageController(messageService)

	server.CustomLogger.SuccessLogger.Info("Starting gRPC server for message service")

	server.startGrpcServer(messageController)
}

func (server *Server) startGrpcServer(jobController *controller.MessageController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		server.CustomLogger.ErrorLogger.Error("Failed to listen in message service: ", listener)
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	messageGW.RegisterMessageServiceServer(grpcServer, jobController)
	if err := grpcServer.Serve(listener); err != nil {
		server.CustomLogger.ErrorLogger.Error("Failed to serve gRPC in message service: ", listener)
		log.Fatalf("failed to serve: %s", err)
	}
}
