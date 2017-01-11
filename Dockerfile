FROM golang:latest

ADD . /go/src/SysnotifsPurge
WORKDIR /go/src/SysnotifsPurge
RUN go get github.com/gorilla/mux

RUN go install SysnotifsPurge

ENTRYPOINT /go/bin/SysnotifsPurge

EXPOSE 8080

#docker build . -t sysnotifspurge
#docker run -d -p 127.0.0.1:8080:8080 sysnotifspurge 