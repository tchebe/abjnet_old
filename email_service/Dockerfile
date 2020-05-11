FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY go.mod . 
COPY go.sum .
#use the ip address of the ubuntu when behind a proxy 
ENV http_proxy 'http://nsia%5Cthibaut.zehi:ZjjTEnt%401988@10.10.4.30:8080' 
ENV https_proxy 'http://nsia%5Cthibaut.zehi:ZjjTEnt%401988@10.10.4.30:8080'
RUN go mod download -x
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o abjnetEmailService

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/abjnetEmailService .

ENTRYPOINT ["./abjnetEmailService"]