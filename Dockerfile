FROM alpine:3.17

# Define the project name | 定义项目名称
ARG PROJECT=mcms
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=mcms.yaml
# Define the author | 定义作者
ARG AUTHOR="yuansu.china.work@gmail.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

ENV TZ=Asia/Shanghai
RUN apk update --no-cache && apk add --no-cache tzdata

COPY ./${PROJECT}_rpc ./
COPY ./etc/${CONFIG_FILE} ./etc/

ENTRYPOINT ./${PROJECT}_rpc -f etc/${CONFIG_FILE}