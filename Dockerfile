FROM golang:1.22-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o uchet ./cmd/main.go

CMD ["./uchet"]