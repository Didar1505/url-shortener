package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Didar1505/url-shortener/api-gateway/client"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type URLHandler struct {
	urlClient *client.URLServiceClient
}

func NewURLHandler(urlClient *client.URLServiceClient) *URLHandler {
	return &URLHandler{
		urlClient: urlClient,
	}
}

type CreateShortURLRequest struct {
	URL string `json:"url" binding:"required"`
}

type CreateShortURLResponse struct {
	ShortCode string `json:"short_code"`
	ShortURL  string `json:"short_url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *URLHandler) CreateShortURL(c *gin.Context) {
	var reqBody CreateShortURLRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "invalid request body: url is required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	resp, err := h.urlClient.CreateShortURL(ctx, reqBody.URL)
	if err != nil {
		handleGRPCError(c, err)
		return
	}

	c.JSON(http.StatusCreated, CreateShortURLResponse{
		ShortCode: resp.ShortCode,
		ShortURL:  resp.ShortUrl,
	})
}

func (h *URLHandler) RedirectToOriginalURL(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "short code is required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	resp, err := h.urlClient.GetOriginalURL(ctx, code)
	if err != nil {
		handleGRPCError(c, err)
		return
	}

	c.Redirect(http.StatusFound, resp.OriginalUrl)
}

func handleGRPCError(c *gin.Context, err error) {
	st, ok := status.FromError(err)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "internal server error",
		})
		return
	}

	switch st.Code() {
	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: st.Message(),
		})
	case codes.NotFound:
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error: st.Message(),
		})
	default:
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: st.Message(),
		})
	}
}