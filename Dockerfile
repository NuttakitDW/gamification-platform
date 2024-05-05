FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . /app

RUN go build -v -mod=readonly -o /go/bin/main ./cmd/api

RUN go install github.com/cosmtrek/air@v1.40.4


EXPOSE 8080

# CMD ["/go/bin/main"]
CMD ["air", "-c", ".air.toml"]



