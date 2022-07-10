package startup

import (
	//"context"
	followGw "common/proto/follow_service"
	messageGw "common/proto/message_service"
	postGw "common/proto/post_service"
	recommendationGw "common/proto/recommendation_service"
	"common/tracer"
	"context"
	"fmt"
	cfg "gateway/startup/config"
	"io"
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"log"
	"net/http"

	jobGw "common/proto/job_service"
	userGw "common/proto/user_service"

	"github.com/urfave/negroni"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Name = "gateway"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
	tracer opentracing.Tracer
	closer io.Closer
}

func NewServer(config *cfg.Config) *Server {
	tracer, closer := tracer.Init(Name)
	otgo.SetGlobalTracer(tracer)

	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
		tracer: tracer,
		closer: closer,
	}

	server.initHandlers()
	return server
}

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

func tracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err == nil || err == opentracing.ErrSpanContextNotFound {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				"ServeHTTP",
				// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			endpointName := r.Method + " " + r.URL.Path
			span := tracer.StartSpanFromRequest(endpointName, serverSpan.Tracer(), r)
			defer span.Finish()

			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		)}

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

	recommendationEndpoint := fmt.Sprintf("%s:%s", server.config.RecommendationHost, server.config.RecommendationPort)
	errRec := recommendationGw.RegisterRecommendationServiceHandlerFromEndpoint(context.TODO(), server.mux, recommendationEndpoint, opts)
	if errRec != nil {
		panic(errRec)
	}

}

func cors(h http.Handler, server *Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")

		if r.Method == "OPTIONS" {
			return
		}
		//h.ServeHTTP(w, r)

		endpointName := r.Method + " " + r.URL.Path

		parentSpanContext, err2 := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err2 == nil || err2 == opentracing.ErrSpanContextNotFound {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				endpointName,
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		lrw := negroni.NewResponseWriter(w)
		server.mux.ServeHTTP(lrw, r)

		statusCode := lrw.Status()
		ipAddr := r.RemoteAddr
		fmt.Println("IP ADDRESS:", ipAddr)
		browser := r.UserAgent()
		fmt.Println("BROWSER:", browser)
		t := time.Now()
		fmt.Println("TIMESTAMP:", t.Format("2006-01-02 15:04:00"))

		visitorLabel := prometheus.Labels{
			"ip":        ipAddr,
			"browser":   browser,
			"timestamp": t.Format("2006-01-02 15:04:05"),
		}
		visitor.With(visitorLabel).Inc()

		gb := r.ContentLength
		fmt.Println(gb)
		dataFlowFromReq.Add(float64(gb))
		fmt.Println(dataFlowFromReq)

		totalReq.Inc()
		if statusCode >= 200 && statusCode <= 399 {
			successfulReq.Inc()
		} else if statusCode >= 400 && statusCode <= 599 {
			if statusCode == 404 {
				labels := prometheus.Labels{
					"code":   "404",
					"method": endpointName,
				}
				notFoundReq.With(labels).Inc()
			}
			failedReq.Add(3)
		}
	})
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), cors(server.mux, server)))
}

func (server *Server) GetTracer() opentracing.Tracer {
	return server.tracer
}

func (server *Server) GetCloser() io.Closer {
	return server.closer
}

func (server *Server) CloseTracer() error {
	return server.closer.Close()
}

var (
	totalReq = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dislinkt_total_req",
		Help: "The total number of requests",
	})
	successfulReq = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dislinkt_successful_req",
		Help: "The number of successful requests",
	})
	failedReq = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dislinkt_failed_req",
		Help: "The total number of failed requests",
	})
	notFoundReq = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "dislinkt_not_found_req",
		Help: "The total number of 404 requests with endpoint",
	}, []string{"code", "method"})
	visitor = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "dislinkt_visitor_req",
		Help: "Visitor from request",
	}, []string{"ip", "browser", "timestamp"})
	dataFlowFromReq = promauto.NewCounter(prometheus.CounterOpts{
		Name: "dislinkt_data_flow_req",
		Help: "Data flow from request",
	})
)
