FROM golang:latest

WORKDIR /go/src/algocook/users

ENV SRC_DIR=/go/src/algocook/users
ADD . ${SRC_DIR}

RUN go get -v -u github.com/algocook/proto/users
RUN go get -v -u google.golang.org/grpc/grpclog

RUN cd ${SRC_DIR}
RUN go build -o main ./cmd/server
CMD ["./main"]