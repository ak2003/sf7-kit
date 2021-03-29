# Build Stage
FROM golang:1.15.2 AS builder

ARG SERVICE_NAME

WORKDIR $GOPATH/src/gt-kit

ADD . .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app $SERVICE_NAME/cmd/main.go

# Final Stage
FROM alpine:latest

ARG SERVICE_NAME
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /root/

RUN mkdir logs
RUN chmod 777 logs

COPY --from=builder /go/src/gt-kit/app .
COPY --from=builder /go/src/gt-kit/$SERVICE_NAME/config/config.json ./$SERVICE_NAME/config/config.json

CMD ["./app"]
EXPOSE 9090