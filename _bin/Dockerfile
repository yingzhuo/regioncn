FROM registry.cn-shanghai.aliyuncs.com/yingzhuo/busybox:1.28.4-wfr

MAINTAINER 应卓 yingzhor@gmail.com

ARG VERSION

COPY ./regioncn-linux-amd64-$VERSION /bin/regioncn

COPY docker-entrypoint.sh /bin/docker-entrypoint

RUN chmod +x /bin/regioncn

ENV \
    PATH=$PATH:/usr/local/bin \
    REGIONCN_VERSION=$VERSION

CMD ["docker-entrypoint"]