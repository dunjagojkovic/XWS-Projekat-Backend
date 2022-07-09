package main

import (
	messageGW "common/proto/message_service"
	saga "common/saga/messaging"
	"common/saga/messaging/nats"
	"common/tracer"
	"context"
	"fmt"
	"io"
	"log"
	"messageS/config"
	"messageS/controller"
	"messageS/repository"
	"messageS/service"
	"net"

	otgo "github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	config       *config.Config
	CustomLogger *controller.CustomLogger
	Tracer       otgo.Tracer
	Closer       io.Closer
}

const (
	QueueGroup = "order_service"
	Name       = "message_service"
)

func NewServer(config *config.Config) *Server {
	CustomLogger := controller.NewCustomLogger()
	tracer, closer := tracer.Init(Name)
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config:       config,
		CustomLogger: CustomLogger,
		Tracer:       tracer,
		Closer:       closer,
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

	commandPublisher := server.initPublisher(server.config.CreateMessageCommandSubject)
	replySubscriber := server.initSubscriber(server.config.CreateMessageReplySubject, QueueGroup)
	createMessageOrchestrator := server.initCreateMessageOrchestrator(commandPublisher, replySubscriber)

	messageService := service.NewMessageService(store, createMessageOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.CreateMessageCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateMessageReplySubject)
	server.initCreateMessageHandler(messageService, replyPublisher, commandSubscriber)

	messageController := controller.NewMessageController(messageService)

	server.CustomLogger.SuccessLogger.Info("Starting gRPC server for message service")

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
