FROM golang:latest

ADD . /go/src/SysnotifsPurge
WORKDIR /go/src/SysnotifsPurge
RUN go get github.com/gorilla/mux
RUN go get github.com/pkg/errors

RUN go install SysnotifsPurge/Config
RUN go install SysnotifsPurge/Elasticsearch
RUN go install SysnotifsPurge/Handlers
RUN go install SysnotifsPurge/Routes
RUN go install SysnotifsPurge

ENTRYPOINT /go/bin/SysnotifsPurge

EXPOSE 8080

#docker build . -t sysnotifspurge
#docker run -d -p 127.0.0.1:8080:8080 -e SYSNOTIFSPURGE_ES_ENDPOINT='https://search-core-sysnotifs-prod-7tlhh72bx4aasz7we2d3cjvybi.eu-west-1.es.amazonaws.com'' -e SYSNOTIFSPURGE_THRESHOLD=30 sysnotifspurge 
