package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"userS/config"
	"userS/controller"
	"userS/repository"
	"userS/service"

	userGW "common/proto/user_service"

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

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.UserDBHost, server.config.UserDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
		server.CustomLogger.ErrorLogger.WithFields(logrus.Fields{"User db host": server.config.UserDBHost, "User db port": server.config.UserDBPort}).Error("MongoDB initialization for user service failed")

	} else {
		fmt.Println(uri)
		server.CustomLogger.SuccessLogger.Info("MongoDB initialization for user service succesfull")
	}
	store := repository.NewUserStore(client)
	userService := service.NewUserService(store)
	userController := controller.NewUserController(userService)

	server.CustomLogger.SuccessLogger.Info("Starting gRPC server for user service")
	server.startGrpcServer(userController)
}

func (server *Server) startGrpcServer(userController *controller.UserController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		server.CustomLogger.ErrorLogger.Error("Failed to listen in user service: ", listener)
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userGW.RegisterUserServiceServer(grpcServer, userController)
	if err := grpcServer.Serve(listener); err != nil {
		server.CustomLogger.ErrorLogger.Error("Failed to serve gRPC in user service: ", listener)
		log.Fatalf("failed to serve: %s", err)
	}
}
