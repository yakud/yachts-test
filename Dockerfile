# # # # # # # # # # # # # # # # # # # #
# Build application
# # # # # # # # # # # # # # # # # # # #
FROM golang:1.12-alpine as builder

RUN apk add --no-cache git mercurial ca-certificates openssh-client bash

# Write ssh key
RUN mkdir /root/.ssh && \
    echo "StrictHostKeyChecking no " > /root/.ssh/config

ENV TARGET_FOLDER=/go/src/github.com/yakud/yachts-test
ENV GOPATH=/go

ADD ./ ${TARGET_FOLDER}
WORKDIR ${TARGET_FOLDER}

RUN go get -v -d ./...
RUN go build -o /rest-server ./cmd/rest-server/main.go

# # # # # # # # # # # # # # # # # # # #
# Run application
# # # # # # # # # # # # # # # # # # # #
FROM alpine:latest

COPY --from=builder /rest-server .
COPY --from=builder /go/src/github.com/yakud/yachts-test/static/ /static

WORKDIR /

CMD ["/rest-server"]
