FROM golang:1.20

ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN apt-get update && apt-get install -y gcc

WORKDIR /app
COPY . .

RUN go build -o app telegram/main.go

CMD ["./app"]
