# Golag
# https://hub.docker.com/_/golang
FROM golang:1 as golang

RUN GO111MODULE=off go get -u github.com/oxequa/realize && \
    go get -u github.com/gorilla/mux  && \
    go get -u github.com/gin-gonic/gin   && \
    go get -u gorm.io/driver/mysql && \
    go get -u gorm.io/gorm && \
    go get -u github.com/gemcook/pagination-go



WORKDIR /go/src/github.com/docker_go_nginx/