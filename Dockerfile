FROM golang:1.23-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./bin/main ./cmd/app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
