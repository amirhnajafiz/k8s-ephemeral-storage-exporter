FROM golang:1.20-alpine

WORKDIR /bin/app

COPY . .

RUN go build -o main

CMD ./main
