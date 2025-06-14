FROM golang:1.24.2-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o travel-go -ldflags="-extldflags=-static" ./cmd

EXPOSE 3002

CMD ["./travel-go"]