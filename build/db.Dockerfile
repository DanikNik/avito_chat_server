FROM postgres

ENV POSTGRES_USER=postgres
ENV DB_NAME=chat
ENV POSTGRES_PASSWORD=postgres

# Custom initialization scripts
COPY ./build/create_db.sh /docker-entrypoint-initdb.d/create_db.sh
COPY ./init/init.sql /schema.sql

RUN chmod +x /docker-entrypoint-initdb.d/create_db.sh