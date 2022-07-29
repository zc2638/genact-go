FROM golang:1.18-alpine as builder
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct
ENV GO111MODULE=on

WORKDIR /work
ADD . .
RUN make build

FROM alpine:3.6
MAINTAINER zc
LABEL maintainer="zc" \
    email="zc2638@qq.com"

ENV TZ="Asia/Shanghai"

COPY --from=builder /usr/local/bin/genact /usr/local/bin/genact

WORKDIR /work
ENTRYPOINT ["genact"]