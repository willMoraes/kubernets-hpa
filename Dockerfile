FROM golang:alpine

WORKDIR /go/src/app

COPY src/hello/ .

RUN go get -d -v

RUN go build -o /go/bin/hello

ENTRYPOINT ["/go/bin/hello"]