package startup

import (
	//"context"
	followGw "common/proto/follow_service"
	messageGw "common/proto/message_service"
	postGw "common/proto/post_service"
	"common/tracer"
	"context"
	"fmt"
	cfg "gateway/startup/config"
	"io"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	otgo "github.com/opentracing/opentracing-go"

	"log"
	"net/http"

	jobGw "common/proto/job_service"
	userGw "common/proto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Name = "gateway"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
	Tracer otgo.Tracer
	Closer io.Closer
}

func NewServer(config *cfg.Config) *Server {
	tracer, closer := tracer.Init(Name)
	otgo.SetGlobalTracer(tracer)
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
		Tracer: tracer,
		Closer: closer,
	}
	server.initHandlers()
	return server
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.Tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.Closer
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
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

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	errUser := userGw.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if errUser != nil {
		panic(errUser)
	}

	followEndpoint := fmt.Sprintf("%s:%s", server.config.FollowHost, server.config.FollowPort)
	errFollow := followGw.RegisterFollowServiceHandlerFromEndpoint(context.TODO(), server.mux, followEndpoint, opts)
	if errFollow != nil {
		panic(errFollow)
	}

	messageEndpoint := fmt.Sprintf("%s:%s", server.config.MessageHost, server.config.MessagePort)
	errmessage := messageGw.RegisterMessageServiceHandlerFromEndpoint(context.TODO(), server.mux, messageEndpoint, opts)
	if errmessage != nil {
		panic(errmessage)
	}

}

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")

		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func (server *Server) Start() {
	gwmux := runtime.NewServeMux()

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: tracingWrapper(gwmux),
	}
	//cors(server.mux)

	log.Fatal(gwServer.ListenAndServe())
}
