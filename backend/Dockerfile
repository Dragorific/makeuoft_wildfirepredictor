FROM golang:1.7-alpine

ADD . /backend

WORKDIR /backend
ENV GOPATH=/backend/src/vendor

CMD ["go","run","main.go"]