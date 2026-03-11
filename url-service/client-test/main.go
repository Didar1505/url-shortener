package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Didar1505/url-shortener/pkg/url_shortener_v1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	
	client := pb.NewURLShortenerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	testCreateShortURL(ctx, client)
}

func testCreateShortURL(ctx context.Context, client pb.URLShortenerServiceClient) {
	fmt.Println("Calling CreateShortURL...")

	createResp, err := client.CreateShortURL(ctx, &pb.CreateShortURLRequest{
		OriginalUrl: "https://example.com",
	})
	if err != nil {
		log.Fatalf("CreateShortURL failed: %v", err)
	}

	fmt.Println("Short code:", createResp.ShortCode)
	fmt.Println("Short URL:", createResp.ShortUrl)

	testGetOriginalURL(ctx, client, createResp.ShortCode)
}

func testGetOriginalURL(ctx context.Context, client pb.URLShortenerServiceClient, code string) {
	fmt.Println("Calling GetOriginalURL...")

	getResp, err := client.GetOriginalURL(ctx, &pb.GetOriginalURLRequest{
		ShortCode: code,
	})
	if err != nil {
		log.Fatalf("GetOriginalURL failed: %v", err)
	}

	fmt.Println("Original URL:", getResp.OriginalUrl)
}