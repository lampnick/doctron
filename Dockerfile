FROM golang:1.15.2-alpine as builder

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

RUN mkdir -p /doctron
COPY . /doctron

RUN cd /doctron && \
    go build && \
    cp -fr doctron /usr/local/bin && \
    chmod +x /usr/local/bin/doctron

FROM lampnick/runtime:chromium-alpine

MAINTAINER lampnick <nick@lampnick.com>
COPY --from=builder  /usr/local/bin/doctron /usr/local/bin/doctron
COPY conf/default.yaml /doctron.yaml
ENTRYPOINT ["doctron", "--config", "/doctron.yaml"]