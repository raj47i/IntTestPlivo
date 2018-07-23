FROM postgres:10-alpine

COPY plivo.sql /docker-entrypoint-initdb.d/10-init.sql