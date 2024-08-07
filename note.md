docker build -t starrail-relic-builder -f Dockerfile.builder .
docker run --rm -it -v /root/workspace/starrail-relic:/workspace --workdir /workspace --entrypoint /bin/bash starrail-relic-builder


docker run -it --rm -v $PWD:/workspace --entrypoint /bin/bash starrail-relic-builder