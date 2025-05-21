FROM golang:alpine AS builder

WORKDIR /app

COPY ./study_service/go.* ./

RUN go mod download

COPY ./study_service/ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd/main.go

EXPOSE 8080

CMD ["./bin/main"]