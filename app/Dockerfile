FROM golang:1.14

WORKDIR /usr/src/app

COPY . /usr/src/app

RUN go mod tidy

RUN go mod vendor

RUN go get github.com/codegangsta/gin

CMD ["gin"]