package scrapper

import (
	"github.com/achievermina/IndeedWebScrapping-API/scrapper/scrapperpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	scrapperpb.UnimplementedJobServiceServer
}




func StartServer() {
	//gRPC
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:50051 %v", err)
	}
	var grpcServer = grpc.NewServer()
	scrapperpb.RegisterJobServiceServer(grpcServer, &Server{})
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("Server starting...")
}
