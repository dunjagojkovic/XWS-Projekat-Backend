package startup

import (
	//"context"
	postGw "common/proto/post_service"
	"context"
	"crypto/tls"
	"fmt"
	cfg "gateway/startup/config"
	"log"
	"net/http"

	jobGw "common/proto/job_service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("Gateway.crt", "Gateway.key")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func (server *Server) initHandlers() {
	/*tlsCredentials, err1 := loadTLSCredentials()
	if err1 != nil {
		log.Fatal("cannot load TLS credentials")
	}*/
	/*tlsCreds, err2 := credentials.NewClientTLSFromFile("Gateway.crt", "")
	if err2 != nil {
		log.Fatalf("No cert found: %v", err2)
	}*/
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(config))}
	postEmdpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	err := postGw.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEmdpoint, opts)
	if err != nil {
		panic(err)
	}
	jobEmdpoint := fmt.Sprintf("%s:%s", server.config.JobHost, server.config.JobPort)
	errJ := jobGw.RegisterJobServiceHandlerFromEndpoint(context.TODO(), server.mux, jobEmdpoint, opts)
	if errJ != nil {
		panic(errJ)
	}

}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")

		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", server.config.Port), "Gateway.crt", "Gateway.key", cors(server.mux)))
}
