package server

import (
	"context"
	"fmt"

	pb "github.com/Didar1505/url-shortener/pkg/url_shortener_v1"
	"github.com/Didar1505/url-shortener/url-service/logic"
	"github.com/Didar1505/url-shortener/url-service/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type URLShortenerServer struct {
	pb.UnimplementedURLShortenerServiceServer
	store store.URLStore
	generator *logic.CodeGenerator
	baseURL string
}

func NewURLShortenerServer(store store.URLStore, generator *logic.CodeGenerator, baseURL string) *URLShortenerServer {
	return &URLShortenerServer{
		store: store,
		generator: generator,
		baseURL: baseURL, 
	}
}

func (s *URLShortenerServer) CreateShortURL(ctx context.Context, req *pb.CreateShortURLRequest) (*pb.CreateShortURLResponse, error) {
	originalURL := req.GetOriginalUrl()
	if originalURL == "" {
		return nil, status.Error(codes.InvalidArgument, "original_url is required")
	}

	if !logic.IsValidURL(originalURL) {
		return nil, status.Error(codes.InvalidArgument, "invalid URL: must be a valid http or https URL")
	}

	const maxAttempts = 5

	for i := 0; i < maxAttempts; i++ {
		code := s.generator.Generate()

		err := s.store.Save(code, originalURL)
		if err == nil {
			return &pb.CreateShortURLResponse{
				ShortCode: code,
				ShortUrl:  fmt.Sprintf("%s/u/%s", s.baseURL, code),
			}, nil
		}

		if err != store.ErrCodeAlreadyExists {
			return nil, status.Errorf(codes.Internal, "failed to save short URL: %v", err)
		}
	}

	return nil, status.Error(codes.Internal, "failed to generate a unique short code after several attempts")
}

func (s *URLShortenerServer) GetOriginalURL(ctx context.Context, req *pb.GetOriginalURLRequest) (*pb.GetOriginalURLResponse, error) {
	shortCode := req.GetShortCode()
	if shortCode == "" {
		return nil, status.Error(codes.InvalidArgument, "short_code is required")
	}

	originalURL, found := s.store.Get(shortCode)
	if !found {
		return nil, status.Error(codes.NotFound, "short code not found")
	}

	return &pb.GetOriginalURLResponse{
		OriginalUrl: originalURL,
	}, nil
}