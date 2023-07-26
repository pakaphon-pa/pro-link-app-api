FROM postgres:12-alpine

COPY db/init-db.sql /docker-entrypoint-initdb.d/