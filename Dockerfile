FROM golang:1.16.5-alpine3.12 as builder
WORKDIR /go/src/github.com/mendersoftware/iot-manager
RUN mkdir -p /etc_extra
RUN echo "nobody:x:65534:" > /etc_extra/group
RUN echo "nobody:!::0:::::" > /etc_extra/shadow
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_extra/passwd
RUN chown -R nobody:nobody /etc_extra
RUN apk add --no-cache \
    xz-dev \
    musl-dev \
    gcc \
    ca-certificates
COPY ./ .
RUN CGO_ENABLED=0 go build

FROM scratch
EXPOSE 8080
COPY --from=builder /etc_extra/ /etc/
USER 65534
WORKDIR /etc/iot-manager
COPY --from=builder --chown=nobody /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=nobody ./config.yaml .
COPY --from=builder --chown=nobody /go/src/github.com/mendersoftware/iot-manager/iot-manager /usr/bin/

ENTRYPOINT ["/usr/bin/iot-manager"]
