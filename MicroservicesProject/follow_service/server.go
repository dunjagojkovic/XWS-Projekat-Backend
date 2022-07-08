package main

import (
	"fmt"
	"followS/config"
	"followS/controller"
	"log"
	"net"
	"strings"

	pb "common/proto/follow_service"
	"followS/repository"
	"followS/service"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

func (server *Server) Start() {
	configuration := server.parseConfiguration()
	driver, err := configuration.NewDriver()
	server.CustomLogger.SuccessLogger.Info("Neo4J initialization for follow service successful, PORT: ", server.config.FollowDBPort)

	if err != nil {
		log.Fatal(err)
	}
	followStore := repository.NewFollowStore(&driver, configuration.Database)
	followService := service.NewFollowService(followStore)
	usersEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	followController := controller.NewFollowController(followService, usersEndpoint)

	server.CustomLogger.SuccessLogger.Info("Starting gRPC server for follow service")
	server.startGrpcServer(followController)
}

func (server *Server) startGrpcServer(followersHandler *controller.FollowController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		server.CustomLogger.ErrorLogger.Error("Failed to listen in follow service: ", listener)
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFollowServiceServer(grpcServer, followersHandler)
	if err := grpcServer.Serve(listener); err != nil {
		server.CustomLogger.ErrorLogger.Error("Failed to serve gRPC in follow service: ", listener)
		log.Fatalf("failed to serve: %s", err)
	}
}

func (server *Server) parseConfiguration() *Neo4jConfiguration {
	database := server.config.FollowDatabase
	if !strings.HasPrefix(server.config.DBNeo4jVersion, "4") {
		database = ""
	}
	return &Neo4jConfiguration{
		Url:      fmt.Sprintf("neo4j://%s:%s", server.config.FollowDBHost, server.config.FollowDBPort),
		Username: server.config.FollowDBUsername,
		Password: server.config.FollowDBPassword,
		Database: database,
	}
}

func (nc *Neo4jConfiguration) NewDriver() (neo4j.Driver, error) {
	return neo4j.NewDriver(nc.Url, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}

func (server *Server) initFollowersStore(driver *neo4j.Driver, dbName string) repository.FollowStoreI {
	return repository.NewFollowStore(driver, dbName)
}
