FROM golang as builder
MAINTAINER chenshanping
WORKDIR /app
ENV GOPROXY https://goproxy.cn
ADD backend /app
RUN rm -f /app/logs/* &&  \
    go build -o server main.go
EXPOSE 8080
ENTRYPOINT ["/app/server"]


