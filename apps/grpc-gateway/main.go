package main

import (
	"flag"
	"fmt"
	"net/http"

	gw "api-gateway/proto"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/gorilla/handlers"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("research-api-svc:19003")
	err := gw.RegisterArticleServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}

	newMux := handlers.CORS(
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://research.weeekend.work", "http://api.weeekend.work", "http://localhost:8080", "http://127.0.0.1:8080"}),
		handlers.AllowedHeaders([]string{"content-type"}),
	)(mux)
	return http.ListenAndServe(":5000", newMux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
