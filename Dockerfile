FROM golang:1.17.10-alpine

ADD . /home

WORKDIR /home

RUN \
       apk add --no-cache bash git make go

ENV GOPATH /go
ENV PATH /go/bin:$PATH

CMD ["make"]

# sudo docker build . --tag v1.0
# sudo docker run --publish 8081:8081 v1.0 -d