set -e

GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./
curl -T starrail-relic.exe "http://192.168.100.241/public-writable/starrail-relic.exe"