package main

import (
	"context"
	"fmt"
	"io"
	"jobS/controller"
	"jobS/repository"
	"jobS/service"
	"log"
	"net"

	jobGW "common/proto/job_service"
	"common/tracer"

	otgo "github.com/opentracing/opentracing-go"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

const (
	Name = "job_service"
)

type Server struct {
	config *Config
	Tracer otgo.Tracer
	Closer io.Closer
}

func NewServer(config *Config) *Server {
	tracer, closer := tracer.Init(Name)
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		Tracer: tracer,
		Closer: closer,
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
