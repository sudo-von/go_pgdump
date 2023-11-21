FROM golang:1.19

WORKDIR /go/src/app

COPY . .

RUN apt-get update && \
  apt-get install -y postgresql-client

RUN go build cmd/main.go

CMD ["./main"]

VOLUME /go/src/app/files