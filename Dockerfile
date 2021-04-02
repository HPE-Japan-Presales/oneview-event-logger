FROM alpine:3.13.4

WORKDIR app
ADD bin/oneview-event-logger-v0.1-linux-amd64.tar.gz .
CMD ["/app/oneview-event-logger"]