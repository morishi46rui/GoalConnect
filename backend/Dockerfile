FROM golang:1.23.1 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY docs ./docs

WORKDIR /app/cmd/app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/cmd/app/main .

EXPOSE 8080
CMD ["./main"]
