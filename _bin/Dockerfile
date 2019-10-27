FROM busybox:1.28.4

MAINTAINER 应卓 yingzhor@gmail.com

ARG VERSION

COPY ./regioncn-linux-amd64-$VERSION /bin/regioncn

COPY docker-entrypoint.sh /bin/docker-entrypoint

RUN chmod a+x /bin/docker-entrypoint \
&&  mkdir -p /usr/local/bin \
&&  wget --no-check-certificate "https://github.com/yingzhuo/wait-for-readiness/releases/download/1.0.0/wfr-linux-amd64-1.0.0" -O /usr/local/bin/wfr \
&&  chmod +x /usr/local/bin/wfr

ENV \
    PATH=$PATH:/usr/local/bin \
    REGIONCN_VERSION=$VERSION

CMD ["docker-entrypoint"]