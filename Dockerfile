FROM node:21-alpine AS build

WORKDIR /app

COPY package*.json ./

RUN npm ci

COPY tailwind.config.js ./

COPY static ./static

COPY templates ./templates

RUN npx tailwindcss -i ./static/css/input.css -o ./static/css/output.css


FROM golang:1.21.2

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./

RUN go mod download

COPY --from=build /app ./

RUN templ generate -path ./templates

COPY . .

RUN go build -o main ./cmd/...

EXPOSE 3000

CMD ["./main"]
