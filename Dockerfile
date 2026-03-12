FROM golang:1.25-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/url-service ./url-service/main.go

FROM alpine:3.23

RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

COPY --from=builder /out/url-service /app/url-service

EXPOSE 50051
USER app

ENTRYPOINT ["/app/url-service"]
