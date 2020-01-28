FROM alpine:latest
MAINTAINER 1907004005@qq.com
RUN apk update \
  &7 apk add --no-cache bash
COPY dytt_server /usr/bin/
ENTRYPOINT [ "dytt_server" ]
EXPOSE 5000
