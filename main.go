package main

import (
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/richguo0615/game-streaming/controllers"
	"github.com/richguo0615/game-streaming/proto"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	grpcServer := grpc.NewServer()
	proto.RegisterGameServer(grpcServer, &controllers.GameController{})
	wrappedGrpc := grpcweb.WrapServer(grpcServer,
		grpcweb.WithWebsockets(true),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
			return true
		}),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
	)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedGrpc.ServeHTTP(resp, req)
	}
	httpServer := http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(http.HandlerFunc(handler), &http2.Server{}),
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
