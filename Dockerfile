# Build base contains things take time to download/build that
# won't change often, to allow routine builds further down in
# the Dockerfile to run quickly.
FROM golang:1.12-alpine as build_base

# Alpine requires some setup to be able to build this project.
RUN apk update && apk add bash coreutils git make protobuf

ENV GO111MODULE=on
WORKDIR /go/src/gaproxy

COPY Makefile.truss .
RUN make -f Makefile.truss truss-install

COPY go.mod .
COPY go.sum .
RUN go mod download


# Builder
FROM build_base AS builder
COPY . .
RUN truss service.proto --svcout .
RUN go build -o /go/bin/gaproxy ./cmd/gaproxy-server


# Runtime image
FROM alpine AS goproxy
RUN apk update && apk add ca-certificates
COPY --from=builder /go/bin/gaproxy /bin/gaproxy
ENTRYPOINT ["/bin/gaproxy"]
