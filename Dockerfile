FROM golang

RUN go get golang.org/x/oauth2
RUN go get github.com/digitalocean/godo
