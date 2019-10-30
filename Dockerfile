FROM golang:latest
RUN mkdir -p /go/src/github.com/pashukhin/pik-arenda-test-task
ADD . /go/src/github.com/pashukhin/pik-arenda-test-task
WORKDIR /go/src/github.com/pashukhin/pik-arenda-test-task
RUN go get -v
