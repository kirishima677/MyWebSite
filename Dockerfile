# Golag
# https://hub.docker.com/_/golang
FROM golang:1 as golang

RUN GO111MODULE=off go get -u github.com/oxequa/realize && \
    go get -u github.com/gorilla/mux  && \
    go get github.com/gin-gonic/gin


WORKDIR /go/src/github.com/docker_go_nginx/