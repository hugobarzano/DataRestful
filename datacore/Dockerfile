FROM ubuntu:latest
ENV APP datacore
ADD $APP /usr/local/bin/$APP
EXPOSE 8080

ENTRYPOINT /usr/local/bin/$APP