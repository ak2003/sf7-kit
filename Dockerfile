# Build Stage
FROM golang:1.15.2 AS builder

#ARG SERVICE_NAME

WORKDIR $GOPATH/src/sf7-kit

ADD . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app main.go

# Final Stage
FROM alpine:latest

#ARG SERVICE_NAME
RUN echo -e "http://nl.alpinelinux.org/alpine/v3.5/main\nhttp://nl.alpinelinux.org/alpine/v3.5/community" > /etc/apk/repositories
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /root/

RUN mkdir logs
RUN chmod 777 logs

COPY --from=builder /go/src/sf7-kit/app .
COPY --from=builder /go/src/sf7-kit/config/config.json ./config/config.json

CMD ["./app"]
EXPOSE 9090