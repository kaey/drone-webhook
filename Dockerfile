FROM alpine:3.5
RUN apk -Uuv add ca-certificates \
  && rm -rf /var/cache/apk/*
ADD drone-webhook /drone-webhook
ENTRYPOINT ["/drone-webhook"]
