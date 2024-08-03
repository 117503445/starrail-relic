FROM golang:1.22.5-bookworm

COPY ./debian.sources /etc/apt/sources.list.d/debian.sources

RUN apt-get update && apt-get install -y gcc-multilib gcc-mingw-w64 libz-mingw-w64-dev

# Bitmap
RUN apt install -y libpng++-dev

RUN go env -w GOPROXY=https://goproxy.cn,direct

ENTRYPOINT ["/workspace/bin.sh"]

# RUN GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./