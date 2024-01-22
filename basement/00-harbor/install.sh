#!/bin/bash

# 执行命令遇到错误就退出
set -e
# 脚本中遇到不存在的变量就退出
# set -u
# 执行指令的时候，同时把指令输出，方便观察结果
set -x
# 执行管道的时候，如果前面的命令出错，管道后面的命令会停止
set -o pipefail


HARBOR_IPADDR=192.168.18.199
HARBOR_PASSWORD=123456
HARBOR_VERSION=v2.8.2

# 在线安装，也就是通过网络下载镜像
wget https://ghproxy.com/https://github.com/goharbor/harbor/releases/download/${HARBOR_VERSION}/harbor-online-installer-${HARBOR_VERSION}.tgz
tar -zxvf harbor-online-installer-${HARBOR_VERSION}.tgz && cd harbor

# 离线安装，也就是镜像在压缩包当中
#wget https://ghproxy.com/https://github.com/goharbor/harbor/releases/download/${HARBOR_VERSION}/harbor-offline-installer-${HARBOR_VERSION}.tgz
#tar -zxvf harbor-offline-installer-${HARBOR_VERSION}.tgz && cd harbor

# 修改hostname
sed -i "s/hostname.*/hostname: ${HARBOR_IPADDR}/g" harbor.yml.tmpl
# 修改Harbor管理员密码
sed -i "s/harbor_admin_password.*/harbor_admin_password: ${HARBOR_PASSWORD}/g" harbor.yml.tmpl
# 注释https
sed -i "s/^https:/#https:/g" harbor.yml.tmpl
sed -i "s/port: 443/#port: 443/g" harbor.yml.tmpl
sed -i "s@certificate: /your/certificate/path@#certificate: /your/certificate/path@g" harbor.yml.tmpl
sed -i "s@private_key: /your/private/key/path@#private_key: /your/private/key/path@g" harbor.yml.tmpl

mv harbor.yml.tmpl harbor.yml
./prepare
./install.sh

docker compose down && docker compose up -d

