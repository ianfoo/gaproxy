# gaproxy

Simple [Truss](https://github.com/metaverse/truss)-generated GoKit service as
part of proof of concept for OAuth based authentication and Google Analytics
Reporting API access.

## Preconditions for building

If you're not going to change the service definition (`service.proto`) or are
always going to build a Docker image, you can skip this section.

If you're going to change the service definition and want to build natively on
your machine, though, you'll need Truss installed so it can generate code for
the service defintions.  First, though, Truss depends on
[`protoc`](https://github.com/protocolbuffers/protobuf/releases/latest)
installed. Download `protoc` for your OS and architecture, or use a package
manager like [Homebrew](https://brew.sh) on macOS.

The Makefile in this repo can install Truss for you.
```
make truss-install
```

This will leave a cloned copy of the Truss git repository on your system,
since there are protobuf-related files in the Truss repo that are required
when Truss runs.

## Building

To build for local run:
```
make build
```

To build Docker image:
```
make docker-build
```

If you want to customize the image tag, set the `IMAGE_TAG` environment
variable, like so:
```
IMAGE_TAG=my_image_name:v0.1.1 make docker-build
```

## Running
To run locally:
```
make run
```

To run in Docker:
```
make docker-run
```
