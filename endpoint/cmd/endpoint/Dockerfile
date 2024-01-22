FROM centos:centos7.9.2009

COPY endpoint /usr/local/bin/endpoint
COPY Endpoint.viper.yaml /etc/endpoint/Endpoint.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

RUN yum install ca-certificates curl -y

WORKDIR /usr/local/bin/

CMD ./endpoint run
