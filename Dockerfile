FROM golang:1.21.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/...

EXPOSE 3000

CMD ["./main"]
