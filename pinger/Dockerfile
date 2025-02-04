FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/app

RUN go build -o pinger .

# Устанавливаем docker-cli и iputils
RUN apk add --no-cache docker-cli iputils

CMD ["./pinger"]