package main

import (
	messageGW "common/proto/message_service"
	saga "common/saga/messaging"
	"common/saga/messaging/nats"
	"context"
	"fmt"
	"log"
	"messageS/controller"
	"messageS/repository"
	"messageS/service"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	config *Config
}

const (
	QueueGroup = "order_service"
)

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.MessageDBHost, server.config.MessageDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uri)
	}
	store := repository.NewMessageStore(client)

	commandPublisher := server.initPublisher(server.config.CreateMessageCommandSubject)
	replySubscriber := server.initSubscriber(server.config.CreateMessageReplySubject, QueueGroup)
	createMessageOrchestrator := server.initCreateMessageOrchestrator(commandPublisher, replySubscriber)

	messageService := service.NewMessageService(store, createMessageOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.CreateMessageCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateMessageReplySubject)
	server.initCreateMessageHandler(messageService, replyPublisher, commandSubscriber)

	messageController := controller.NewMessageController(messageService)

	server.startGrpcServer(messageController)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initCreateMessageOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *service.CreateMessageOrchestrator {
	orchestrator, err := service.NewCreateMessageOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initCreateMessageHandler(service *service.MessageService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := controller.NewCreateMessageCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) startGrpcServer(jobController *controller.MessageController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	messageGW.RegisterMessageServiceServer(grpcServer, jobController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
