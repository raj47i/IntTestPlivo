FROM alpine:latest
LABEL maintainer="Raj <web@arunraj.in>"

RUN adduser -D -G users plivo && \
    apk update && \
    apk add --no-cache openssl ca-certificates wget && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/* && \
    chown plivo:users -Rf /home/plivo/

COPY ./app /home/plivo/app
COPY ./config-docker.json /home/plivo/config.json
RUN chmod +x /home/plivo/app

WORKDIR /home/plivo/

EXPOSE 8080
USER plivo
ENTRYPOINT ["/home/plivo/app"]