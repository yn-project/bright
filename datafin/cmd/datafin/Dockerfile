FROM debian:stable-slim

COPY datafin /usr/local/bin/datafin
COPY DataFin.viper.yaml /etc/datafin/DataFin.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR


RUN apt  update \
    && apt install -y --no-install-recommends ca-certificates curl

WORKDIR /usr/local/bin/

CMD ./datafin run