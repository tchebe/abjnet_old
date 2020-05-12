FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on


COPY go.mod . 
COPY go.sum .
#use the ip address of the ubuntu when behind a proxy 
# ENV http_proxy 'http://10.11.101.77:3128/' 
# ENV https_proxy 'http://10.11.101.77:3128/'
RUN go mod download -x
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o abjnetTaskrunnerService

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/abjnetTaskrunnerService .

CMD ["./abjnetTaskrunnerService"]