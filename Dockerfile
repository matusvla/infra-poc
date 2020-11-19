FROM golang:1.15

WORKDIR /go/src/app
COPY ./go .

RUN go get -d -v ./...
RUN go build -v ./...

CMD ./app

#TODO two level build