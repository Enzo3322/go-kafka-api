FROM golang:1.21.0

WORKDIR /go/src
RUN apt update && apt install build-essential librdkafka-dev && go get -u -y
CMD ["tail", "-f", "/dev/null"]