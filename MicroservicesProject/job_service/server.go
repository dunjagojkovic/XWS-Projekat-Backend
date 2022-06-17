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

func accessibleRoles() map[string][]string {
	const servicePath = "/job.JobService/"
	return map[string][]string{
		/*servicePath + "GetAll":         {"GetAll"},
		servicePath + "OwnerJobOffers": {"OwnerJobOffers"},
		servicePath + "CreateJobOffer": {"CreateJobOffer"},
		servicePath + "AddKey":         {"AddKey"},
		servicePath + "JobOfferSearch": {"JobOfferSearch"},*/
	}
}

func permissionsOfRoles() map[string][]string {
	return map[string][]string{
		"user":  {"GetAll", "OwnerJobOffers", "CreateJobOffer", "AddKey", "JobOfferSearch"},
		"admin": {},
	}
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
	interceptor := NewAuthInterceptor(accessibleRoles(), permissionsOfRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	jobGW.RegisterJobServiceServer(grpcServer, jobController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
