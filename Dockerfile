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

RUN mkdir /app
WORKDIR /app/

RUN mkdir logs
RUN chmod 777 logs

COPY --from=builder /go/src/sf7-kit/app .
COPY --from=builder /go/src/sf7-kit/config/config.json ./config/config.json

CMD ["./app"]
EXPOSE 9090


#FROM golang:1.15.2
#RUN mkdir /app
#ADD . /app
#WORKDIR /app
### Add this go mod download command to pull in any dependencies
#RUN go mod vendor
### Our project will now successfully build with the necessary go libraries included.
#RUN go build -o main .
### Our start command which kicks off
### our newly created binary executable
#CMD ["/app/main"]
#EXPOSE 9090 7000