FROM golang:1.24.2-alpine

RUN mkdir -p /app

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod download

EXPOSE 8080

CMD ["air","-c",".air.toml"]

