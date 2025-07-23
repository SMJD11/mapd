VERSION 0.7
FROM golang:1.24-alpine3.21
WORKDIR /mapd

deps:
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

build:
    FROM +deps
    COPY *.go .
    COPY *.json .
    RUN CGO_ENABLED=0 go build -ldflags="-extldflags=-static -s -w" -o build/mapd
    SAVE ARTIFACT build/mapd /mapd AS LOCAL build/mapd

test-deps:
    FROM +deps
    RUN mkdir .snapshots
    COPY .snapshots/* .snapshots/
    COPY *.go .
    COPY *.json .

test-runner:
    FROM +test-deps
    RUN go test .

test:
    BUILD --platform=linux/arm64 +test-runner

update-snapshots-runner:
    FROM +test-deps
    RUN UPDATE_SNAPSHOTS=true go test . || echo "Snapshot changes generated"
    SAVE ARTIFACT .snapshots/* AS LOCAL .snapshots/

update-snapshots:
    BUILD --platform=linux/arm64 +update-snapshots-runner

format-deps:
    FROM +deps
    RUN go install mvdan.cc/gofumpt@latest

format:
    FROM +format-deps
    COPY *.go .
    COPY *.json .
    RUN gofumpt -l -w .
    SAVE ARTIFACT ./*.go AS LOCAL ./

lint-deps:
    FROM +format-deps
    RUN go install honnef.co/go/tools/cmd/staticcheck@latest

lint:
    FROM +lint-deps
    COPY *.go .
    COPY *.json .
    RUN staticcheck -f stylish .
    RUN test -z $(gofumpt -l -d .)


capnp-deps:
    RUN apk add capnproto-dev
    RUN apk add git
    RUN go install capnproto.org/go/capnp/v3/capnpc-go@latest
    RUN git clone https://github.com/capnproto/go-capnp ../go-capnp

compile-capnp:
    FROM +capnp-deps
    COPY *.capnp .
    RUN capnp compile -I ../go-capnp/std -ogo offline.capnp
    SAVE ARTIFACT offline.capnp.go /offline.capnp.go AS LOCAL offline.capnp.go

build-release:
    BUILD --platform=linux/arm64 +build

docker:
    FROM ubuntu:latest
    WORKDIR /app
    COPY +build/mapd .
    COPY scripts/*.sh .
    RUN apt update
    RUN apt install rclone wget osmium-tool -y
    CMD ["./docker_entry.sh"]
    SAVE IMAGE --push pfeiferj/openpilot-mapd:latest

docker-all-archs:
    BUILD --platform=linux/arm64 +docker
    BUILD --platform=linux/amd64 +docker

build-local:
    FROM +deps
    ARG GOOS
    ARG GOARCH
    COPY *.go .
    COPY *.json .
    RUN CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-extldflags=-static -s -w" -o build/mapd
    SAVE ARTIFACT build/mapd /mapd AS LOCAL build/mapd
