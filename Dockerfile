FROM centos:7

ENV LANG en_US.UTF-8
ENV TZ Asia/Shanghai

WORKDIR /data

COPY  ./redis-web-manager .
COPY ./yaml ./yaml
CMD ["/data/redis-web-manager"]