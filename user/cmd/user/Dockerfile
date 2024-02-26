FROM centos:centos7.9.2009

COPY user /usr/local/bin/user
COPY User.viper.yaml /etc/User/User.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR

WORKDIR /usr/local/bin/

CMD ./user run