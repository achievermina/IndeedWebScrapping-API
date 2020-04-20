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
	var grpcServer = grpc.NewServer()
	scrapperpb.RegisterJobServiceServer(grpcServer, &Server{})
	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		log.Fatalf("could not listen to 0.0.0.0:3000 %v", err)
	}
	log.Println("Server starting...")
	log.Fatal(grpcServer.Serve(listen))
}
