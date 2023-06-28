
FROM golang:1.19

ENV OS "darwin linux windows"
ENV ARCH "amd64"

COPY . ${GOPATH}/src/github.com/AliyunContainerService/docker-machine-driver-aliyunecs
WORKDIR ${GOPATH}/src/github.com/AliyunContainerService/docker-machine-driver-aliyunecs
RUN set -ex \
	&& uname -a \
	&& go version \
	&& go env \
    && go mod tidy \
    && go vet ./...
RUN dmver=v0.13.0 \
    && echo "VERSION docker-machine '$dmver'"
RUN for GOOS in $OS; do \
        arch="$GOOS-$ARCH" \
        && binary="bin/docker-machine-driver-aliyunecs.$arch" \
        && echo "Building $binary" \
        && GOOS=$GOOS GOARCH=$GOARCH go build -o $binary \
        && tar czvf bin/docker-machine-driver-aliyunecs_$arch.tgz $binary; \
	done
