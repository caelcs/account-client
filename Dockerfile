FROM golang:1.12.6-alpine

# Install git
RUN set -ex; \
    apk update; \
    apk add --no-cache git

RUN go get github.com/stretchr/testify

# Set working directory
RUN mkdir /project
ADD . /project
WORKDIR /project

# Run tests
CMD CGO_ENABLED=0 go test ./...