#!/bin/bash
MY_FOLDER=$(cd "$(dirname "$0")";pwd)

rm ${MY_FOLDER}/.traefik -rf
cp -r ${MY_FOLDER}/../traefik ${MY_FOLDER}/.traefik

# git clone https://github.com/NpoolPlatform/traefik.git .traefik; cd .traefik; git checkout entropy-v2.5.3

cd ${MY_FOLDER}/
cp Makefile.service .traefik/Makefile
cp build.Dockerfile.service .traefik/build.Dockerfile

# cd ${MY_FOLDER}/.traefik; mkdir -p v2; cp * v2 -rf | true; rm -rf v2/v2; make generate-crd

cd ${MY_FOLDER}/.traefik; make traefik-binary
mkdir -p .traefik-release

cp ${MY_FOLDER}/.traefik/dist/traefik ${MY_FOLDER}/.traefik/.traefik-release
cp ${MY_FOLDER}/entrypoint.sh ${MY_FOLDER}/.traefik/.traefik-release
cp ${MY_FOLDER}/.traefik/script/ca-certificates.crt ${MY_FOLDER}/.traefik/.traefik-release
cp ${MY_FOLDER}/Dockerfile.service ${MY_FOLDER}/.traefik/.traefik-release/Dockerfile
        

set +e
docker images | grep entropypool | grep traefik-service
rc=$?
set -e
if [ 0 -eq $rc ]; then
docker rmi bright/traefik-service:v2.5.3.6 | true
fi
cd ${MY_FOLDER}/.traefik/.traefik-release; docker build -t bright/traefik-service:v2.5.3.6 .
