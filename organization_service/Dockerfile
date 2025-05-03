FROM golang:alpine AS builder

WORKDIR /app

COPY ./organization_service/go.* ./

RUN go mod download

COPY ./organization_service/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/main.go

EXPOSE 8081

CMD ["./bin/main"]