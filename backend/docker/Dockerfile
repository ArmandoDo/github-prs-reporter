### Stage 1 - Build
###
### Build executable from code
FROM golang:1.21.6-alpine as build
WORKDIR /app

### Set up environment
RUN apk update --no-cache \
  && apk add --no-cache \
    gcc \
    musl-dev \
    make

### Build server app in bin folder
COPY . .
RUN go mod download; \
    make build

### Stage 2 - Install
###
### Build customize docker image for go app
FROM alpine:3.18

### Copy executable
COPY --from=build /app/bin /app/bin

### Start service using command
CMD ["/app/bin/server"]