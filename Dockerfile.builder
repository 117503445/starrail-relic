FROM golang:1.22.5-bookworm

# in china
# COPY ./scripts/debian.sources /etc/apt/sources.list.d/debian.sources
# RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN apt-get update && apt-get install -y gcc-multilib gcc-mingw-w64 libz-mingw-w64-dev g++-mingw-w64 libpng++-dev

# if code is belongs to non-root user, add this to avoid permission issue
RUN git config --global --add safe.directory /workspace

WORKDIR /workspace

ENTRYPOINT ["/workspace/scripts/bin.sh"]