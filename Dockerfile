FROM golang:latest

WORKDIR /go/src/algocook/users

ENV SRC_DIR=/go/src/algocook/users

COPY go.mod ${SRC_DIR}
COPY go.sum ${SRC_DIR}
RUN go mod download

ADD . ${SRC_DIR}
RUN cd ${SRC_DIR}
RUN go build -o main ./cmd/server
CMD ["./main"]