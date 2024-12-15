FROM ubuntu:jammy

COPY endpoint /usr/local/bin/endpoint
COPY Endpoint.viper.yaml /etc/endpoint/Endpoint.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

WORKDIR /usr/local/bin/

CMD ./endpoint run
