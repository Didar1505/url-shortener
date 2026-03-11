package server

import (
	"context"
	"testing"

	pb "github.com/Didar1505/url-shortener/pkg/url_shortener_v1"
	"github.com/Didar1505/url-shortener/url-service/logic"
	"github.com/Didar1505/url-shortener/url-service/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateShortURL_Success(t *testing.T) {
	urlStore := store.NewMemoryStore()
	generator := logic.NewCodeGenerator(6)
	srv := NewURLShortenerServer(urlStore, generator, "http://localhost:8080")

	resp, err := srv.CreateShortURL(context.Background(), &pb.CreateShortURLRequest{
		OriginalUrl: "https://example.com",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.ShortCode == "" {
		t.Fatal("expected short code to be generated")
	}

	if resp.ShortUrl == "" {
		t.Fatal("expected short URL to be returned")
	}
}

func TestCreateShortURL_InvalidURL(t *testing.T) {
	urlStore := store.NewMemoryStore()
	generator := logic.NewCodeGenerator(6)
	srv := NewURLShortenerServer(urlStore, generator, "http://localhost:8080")

	_, err := srv.CreateShortURL(context.Background(), &pb.CreateShortURLRequest{
		OriginalUrl: "not-a-valid-url",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatal("expected gRPC status error")
	}

	if st.Code() != codes.InvalidArgument {
		t.Fatalf("expected InvalidArgument, got %v", st.Code())
	}
}

func TestGetOriginalURL_Success(t *testing.T) {
	urlStore := store.NewMemoryStore()
	err := urlStore.Save("abc123", "https://example.com")
	if err != nil {
		t.Fatalf("failed to seed store: %v", err)
	}

	generator := logic.NewCodeGenerator(6)
	srv := NewURLShortenerServer(urlStore, generator, "http://localhost:8080")

	resp, err := srv.GetOriginalURL(context.Background(), &pb.GetOriginalURLRequest{
		ShortCode: "abc123",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.OriginalUrl != "https://example.com" {
		t.Fatalf("expected original URL to match, got %q", resp.OriginalUrl)
	}
}

func TestGetOriginalURL_NotFound(t *testing.T) {
	urlStore := store.NewMemoryStore()
	generator := logic.NewCodeGenerator(6)
	srv := NewURLShortenerServer(urlStore, generator, "http://localhost:8080")

	_, err := srv.GetOriginalURL(context.Background(), &pb.GetOriginalURLRequest{
		ShortCode: "missing",
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatal("expected gRPC status error")
	}

	if st.Code() != codes.NotFound {
		t.Fatalf("expected NotFound, got %v", st.Code())
	}
}