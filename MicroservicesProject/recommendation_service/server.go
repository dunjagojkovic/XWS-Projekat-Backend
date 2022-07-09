package main

import (
	"fmt"

	pb "common/proto/recommendation_service"
	"log"
	"net"
	"recommendationS/config"
	"recommendationS/controller"
	"recommendationS/repository"
	"recommendationS/service"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

func (server *Server) Start() {
	neo4jClient := server.initNeo4J()

	recommendationStore := repository.NewRecommendationStore(neo4jClient)
	recommendationService := service.NewRecommendationService(recommendationStore)
	recommendationController := controller.NewRecommendationController(recommendationService)

	server.startGrpcServer(recommendationController)
}

func GetClient(uri, username, password string) (*neo4j.Driver, error) {

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &driver, nil
}

func (server *Server) initNeo4J() *neo4j.Driver {

	//uri := "bolt:\\" + server.config.ConnectionDBHost + ":" + server.config.ConnectionDBPort
	// dbUri := "bolt://localhost:7687"
	dbUri := "bolt://" + server.config.RecommendationDBHost + ":" + server.config.RecommendationDBPort

	client, err := GetClient(dbUri, server.config.RecommendationDBUsername, server.config.RecommendationDBPassword)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) startGrpcServer(recommendationController *controller.RecommendationController) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRecommendationServiceServer(grpcServer, recommendationController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
