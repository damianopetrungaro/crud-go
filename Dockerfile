FROM golang:latest

WORKDIR /go/src/github.damianopetrungaro.testing-golang
COPY . /go/src/github.damianopetrungaro.testing-golang

RUN curl https://glide.sh/get | sh \
    && glide install \
    && go build -o ./executable ./src

EXPOSE 80

ENTRYPOINT ["./executable"]