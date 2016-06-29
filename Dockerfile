FROM golang:1.7

COPY . /go/src/github.com/cultureamp/aws-nquire

RUN go get -u github.com/aws/aws-sdk-go

RUN go build github.com/cultureamp/aws-nquire

RUN cp ./aws-nquire /usr/local/bin

CMD aws-nquire
