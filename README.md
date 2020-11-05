# drone-docker

[![Build Status](http://cloud.drone.io/api/badges/drone-plugins/drone-docker/status.svg)](https://drone.tripanels.com/Sean/drone-docker/)
Drone plugin to build and publish Docker images to a container registry. For the usage information and a listing of the available options please take a look at [the docs](https://wordpress-3238-blog.tripanels.com/drone-docker%e6%b7%bb%e5%8a%a0tag%e7%ae%a1%e7%90%86/).

## Build

Build the binaries with the following commands:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/linux/amd64/drone-docker ./cmd/drone-docker
```
jjj
## Docker

Build the Docker images with the following commands:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/docker/Dockerfile.linux.amd64 --tag plugins/docker .

```

## Usage

> Notice: Be aware that the Docker plugin currently requires privileged capabilities, otherwise the integrated Docker daemon is not able to start.

```console
docker run --rm \
  -e PLUGIN_TAG=latest \
  -e PLUGIN_REPO=octocat/hello-world \
  -e DRONE_COMMIT_SHA=d8dbe4d94f15fe89232e0402c6e8a0ddf21af3ab \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  --privileged \
  plugins/docker --dry-run
```
