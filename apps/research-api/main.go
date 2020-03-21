package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"research-api/adapter/controller"
	"research-api/config"
	"research-api/infrastructure/server"
)

var port = ":19003"

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	c := config.Config{}
	articleController := controller.NewArticleController(&c)
	server.NewServer(s, articleController)

	fmt.Printf("[server started] localhost%s\n", port)
	log.Fatalln(s.Serve(listenPort))
}
