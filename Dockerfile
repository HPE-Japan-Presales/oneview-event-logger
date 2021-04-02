ARG App=oneview-event-logger


FROM golang:alpine3.13 as builder
MAINTAINER taku.kimura@takos-lab.com

RUN apk update && apk add git make
RUN git clone https://github.com/fideltak/oneview-event-logger.git
RUN cd oneview-event-logger && make build
RUN cd oneview-event-logger/bin && tar xvfz oneview-event-logger-$(git describe --tags --abbrev=0)-linux-amd64.tar.gz

FROM alpine:3.13.4
MAINTAINER taku.kimura@takos-lab.com
WORKDIR app
COPY --from=builder /go/oneview-event-logger/bin/oneview-event-logger .
CMD ["/app/oneview-event-logger"]