FROM golang:1.10.4 AS build1

ENV PROJECT_DIR /opt/echo-env-sample

WORKDIR ${PROJECT_DIR}

ADD . ${PROJECT_DIR}

RUN apt-get update -y && \
    apt-get install -y make && \
    make build && \
    chmod a+x echo-env-sample

RUN make build

FROM alpine:3.8

COPY --from=build1 /opt/echo-env-sample/ /opt/echo-env-sample

WORKDIR /opt/echo-env-sample

CMD ["./echo-env-sample"]
