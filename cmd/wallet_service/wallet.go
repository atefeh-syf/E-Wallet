package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//pb "github.com/atefeh-syf/yumigo/api/v1/pb/wallet"
	"github.com/atefeh-syf/yumigo/pkg/wallet"
	"github.com/atefeh-syf/yumigo/pkg/wallet/endpoints"
	"github.com/atefeh-syf/yumigo/pkg/wallet/transport"
	//kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	//"github.com/oklog/oklog/pkg/group"
	//"google.golang.org/grpc"
)

const (
	defaultHTTPPort = "8081"
	defaultGRPCPort = "8082"
)

func main() {
	var (
		logger   log.Logger
		httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort))
		//grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
	)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var (
		service     = wallet.NewWalletService()
		eps         = endpoints.NewEndpointSet(service)
		httpHandler = transport.NewHTTPHandler(eps)
		//grpcServer  = transport.NewGRPCServer(eps)
	)

	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}
	logger.Log("transport", "HTTP", "addr", httpAddr)
	http.Serve(httpListener, httpHandler)

	// grpcListener, err := net.Listen("tcp", grpcAddr)
	// if err != nil {
	// 	logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	// 	os.Exit(1)
	// }

	// logger.Log("transport", "gRPC", "addr", grpcAddr)
	// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
	// the here demonstrated zipkin tracing middleware.
	// baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	// pb.RegisterWalletServer(baseServer, grpcServer)
	// baseServer.Serve(grpcListener)


	// This function just sits and waits for ctrl-C.
	cancelInterrupt := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-c:
		fmt.Errorf("received signal %s", sig)
	case <-cancelInterrupt:
	}

}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}