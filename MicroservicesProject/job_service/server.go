package main

import (
	"context"
	"fmt"
	"jobS/controller"
	"jobS/repository"
	"jobS/service"
	"log"
	"net"

	jobGW "common/proto/job_service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Server struct {
	config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	uri := fmt.Sprintf("mongodb://%s:%s", server.config.JobDBHost, server.config.JobDBPort)
	options := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), options)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uri)
	}
	store := repository.NewJobStore(client)
	jobService := service.NewJobService(store)

	jobController := controller.NewJobController(jobService)

	server.startGrpcServer(jobController)
}

func (server *Server) startGrpcServer(jobController *controller.JobController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	jobGW.RegisterJobServiceServer(grpcServer, jobController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
