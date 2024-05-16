FROM debian:stable-slim

RUN sed -i ‘s@//.*archive.ubuntu.com@//mirrors.aliyun.com@g’ /etc/apt/sources.list
RUN apt-get update
COPY contract /usr/local/bin/contract
COPY Contract.viper.yaml /etc/contract/Contract.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR


RUN apt  update \
    && apt install -y --no-install-recommends ca-certificates curl
RUN apt install gcc-multilib g++-multilib -y

WORKDIR /usr/local/bin/

CMD ./contract run