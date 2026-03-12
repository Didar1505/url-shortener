package main

import (
	"log"
	"os"

	"github.com/Didar1505/url-shortener/api-gateway/client"
	"github.com/Didar1505/url-shortener/api-gateway/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	urlServiceAddr := os.Getenv("URL_SERVICE_ADDR")
	if urlServiceAddr == "" {
		urlServiceAddr = "dprogger.online:50051"
	}

	conn, err := grpc.Dial(urlServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	urlServiceClient := client.NewURLServiceClient(conn)
	urlHandler := handler.NewURLHandler(urlServiceClient)

	router := gin.Default()

	router.POST("/shorten", urlHandler.CreateShortURL)
	router.GET("/u/:code", urlHandler.RedirectToOriginalURL)

	log.Printf("API Gateway is running on :8080 (gRPC target: %s)", urlServiceAddr)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run API gateway: %v", err)
	}
}
