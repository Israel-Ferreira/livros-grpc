FROM golang:1.18

WORKDIR /usr/app

COPY go.sum .
COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o bin/livros-grpc


EXPOSE 9000

CMD ["bin/livros-grpc"]