FROM alpine:3.7

EXPOSE 8080

ENV APP_ROOT=/opt/app-root \
    APP_BIN=/opt/app-root/bin \
    PATH=${APP_BIN}:$PATH \
    TZ='Asia/Shanghai'

RUN  mkdir -p ${APP_BIN} ${APP_ROOT} \
     && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
     && apk update \
     && apk upgrade \
     && apk --no-cache add ca-certificates iputils\
     && apk add -U tzdata ttf-dejavu busybox-extras curl bash\
     && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
     && adduser -u 1001 -S -G root -g 0 -D -h ${APP_ROOT} -s /sbin/nologin go

COPY ./sidecar-injector ${APP_BIN}
# Drop the root user and make the content of /opt/app-root owned by user 1001
RUN chown -R 1001:0 ${APP_ROOT}

RUN chmod 755 ${APP_BIN}/sidecar-injector

WORKDIR ${APP_BIN}

USER 0

ENTRYPOINT ["sidecar-injector"]