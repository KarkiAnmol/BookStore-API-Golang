FROM golang:1.18.8-alpine3.16

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8083

CMD ["go", "run", "cmd/main.go"]