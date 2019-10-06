FROM busybox:1.28.4

MAINTAINER 应卓 yingzhor@gmail.com

ARG VERSION

COPY ./regioncn-linux-amd64-$VERSION /bin/regioncn

ENV REGIONCN_VERSION=$VERSION

CMD ["regioncn"]
