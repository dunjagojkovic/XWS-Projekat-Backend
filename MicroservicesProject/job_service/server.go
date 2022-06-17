package main

import (
	"context"
	"crypto/tls"
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
	"google.golang.org/grpc/credentials"
)

type Server struct {
	config *Config
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("JobServiceBackend.crt", "JobServiceBackend.key")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
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
		"User":  {"GetAll", "OwnerJobOffers", "CreateJobOffer", "AddKey", "JobOfferSearch"},
		"Admin": {},
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

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials")
	}
	interceptor := NewAuthInterceptor(accessibleRoles(), permissionsOfRoles())
	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	jobGW.RegisterJobServiceServer(grpcServer, jobController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
