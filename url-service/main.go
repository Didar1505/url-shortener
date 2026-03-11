package main

import (
	"log"
	"net"

	pb "github.com/Didar1505/url-shortener/pkg/url_shortener_v1"
	"github.com/Didar1505/url-shortener/url-service/logic"
	"github.com/Didar1505/url-shortener/url-service/server"
	"github.com/Didar1505/url-shortener/url-service/store"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	urlStore := store.NewMemoryStore()
	codeGenerator := logic.NewCodeGenerator(6)
	urlServer := server.NewURLShortenerServer(urlStore, codeGenerator, "http://localhost:8080")

	grpcServer := grpc.NewServer()
	pb.RegisterURLShortenerServiceServer(grpcServer, urlServer)

	log.Println("gRPC URL service is running on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}