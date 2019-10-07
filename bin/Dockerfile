FROM quay.io/yingzhuo/my-alpine:3.10.2

MAINTAINER 应卓 yingzhor@gmail.com

ARG VERSION

COPY ./regioncn-linux-amd64-$VERSION /bin/regioncn

COPY docker-entrypoint.sh /bin/docker-entrypoint

RUN chmod +x /bin/docker-entrypoint

ENV \
    REGIONCN_VERSION=$VERSION

CMD ["docker-entrypoint"]
