FROM ubuntu:latest
ENV APP datacore
ADD ./build/$APP /usr/local/bin/$APP
EXPOSE 8080

ENTRYPOINT /usr/local/bin/$APP
