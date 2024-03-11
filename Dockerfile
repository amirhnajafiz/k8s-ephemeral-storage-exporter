FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY internal/ internal
COPY main.go main.go

RUN go build -o main
RUN chmod +x ./main

CMD "./main"
