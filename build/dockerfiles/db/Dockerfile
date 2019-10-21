FROM mysql:8.0.18

ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=users

ADD schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306

