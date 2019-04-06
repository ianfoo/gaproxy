# Go module base:
# Keep modules in their own image to avoid downloading every time.
FROM golang:1.12-alpine as module_base

# Alpine requires some setup to be able to pull Go modules.
RUN apk add bash ca-certificates git
# gcc g++ libc-dev

ENV GO111MODULE=on
WORKDIR /go/src/gaproxy

COPY go.mod .
COPY go.sum .
RUN go mod download


# Builder
FROM module_base AS builder
ENV GO111MODULE=on
COPY . .
RUN go build -o /go/bin/gaproxy ./cmd/gaproxy-server


# Runtime image
FROM alpine AS goproxy
RUN apk update && apk add ca-certificates
COPY --from=builder /go/bin/gaproxy /bin/gaproxy
CMD ["/bin/gaproxy"]
