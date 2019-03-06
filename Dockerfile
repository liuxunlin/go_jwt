FROM golang:1.11

WORKDIR $GOPATH/src/go_wechat
COPY . $GOPATH/src/go_wechat

RUN go build .

EXPOSE 8080
ENTRYPOINT ["./go_wechat"]