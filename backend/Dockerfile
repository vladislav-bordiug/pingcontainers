FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/app

RUN go build -o backend .

EXPOSE ${PORT}
CMD ["./backend"]