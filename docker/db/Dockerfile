#docker build -t db docker/db/ - create image or docker build -f ./docker/db/Dockerfile -t db .
FROM postgres:latest

COPY ./book_server/pkg/db/script/init.sql /docker-entrypoint-initdb.d/init.sql

EXPOSE 5432