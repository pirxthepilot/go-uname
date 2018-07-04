FROM golang:1.10.3-alpine3.7

WORKDIR /go/src/github.com/pirxthepilot/go-uname
COPY . .

ENV GOPATH=/go

RUN go build
CMD ["tail", "-f", "/dev/null"]
