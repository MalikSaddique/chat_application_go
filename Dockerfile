FROM golang:1.24.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 go build -o chatting-app main.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY .env .  
COPY --from=builder /app/chatting-app .
EXPOSE 8003
CMD ["./chatting-app"]