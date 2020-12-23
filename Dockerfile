FROM golang:1.14.4

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct
COPY . /usr/src/alertmanger-feishu-webhook
RUN   cd /usr/src/alertmanger-feishu-webhook \
      && go build

FROM centos:7.6.1810

ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

WORKDIR /root/
COPY --from=0 /usr/src/alertmanger-feishu-webhook/alertmanger-feishu-webhook .
CMD ["./alertmanger-feishu-webhook"]
