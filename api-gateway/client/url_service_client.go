package client

import (
	"context"

	pb "github.com/Didar1505/url-shortener/pkg/url_shortener_v1"

	"google.golang.org/grpc"
)

type URLServiceClient struct {
	client pb.URLShortenerServiceClient
}

func NewURLServiceClient(conn *grpc.ClientConn) *URLServiceClient {
	return &URLServiceClient{
		client: pb.NewURLShortenerServiceClient(conn),
	}
}

func (c *URLServiceClient) CreateShortURL(ctx context.Context, originalURL string) (*pb.CreateShortURLResponse, error) {
	return c.client.CreateShortURL(ctx, &pb.CreateShortURLRequest{
		OriginalUrl: originalURL,
	})
}

func (c *URLServiceClient) GetOriginalURL(ctx context.Context, shortCode string) (*pb.GetOriginalURLResponse, error) {
	return c.client.GetOriginalURL(ctx, &pb.GetOriginalURLRequest{
		ShortCode: shortCode,
	})
}