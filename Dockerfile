FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

RUN go build -v -mod=readonly -o main ./cmd/api

EXPOSE 8080

CMD ["./main"]

