FROM golang:1.25-alpine AS builder

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-w -s" -o ./bin/main ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /src/bin/main .
COPY /.env .

EXPOSE 8080

CMD ["./main"]
