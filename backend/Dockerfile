FROM golang:1.19

LABEL maintainer="Alcides Souza <alcides.braga@gmail.com>"

ARG USER_GID=1000
ARG USER_UID=1000

RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get -y upgrade \
    && DEBIAN_FRONTEND=noninteractive apt-get install \
    curl \
    locales \
    postgresql-client \
    && apt-get -y autoclean autoremove \
    && rm -rf /var/lib/apt/lists/*

RUN echo "en_US.UTF-8 UTF-8" > /etc/locale.gen \
    && locale-gen

# Create a non-privileged user for the application.
RUN groupadd --non-unique --gid $USER_GID app \
    && useradd --non-unique --shell /bin/bash --gid app --uid $USER_UID --home-dir /app app

# Set ownership of files to the app user and set default permissions.
RUN find /app /go | xargs chown $USER_UID:$USER_GID
RUN find /app /go -type f | xargs chmod 0640
RUN find /app /go -type d | xargs chmod 0750

RUN chmod gu+x /app/*-entrypoint.sh

USER app

ENV CGO_ENABLED=0
ENV PATH=/app/bin:$PATH

# Install godotenv: https://github.com/joho/godotenv
RUN cd /go \
    && go install github.com/joho/godotenv/cmd/godotenv@latest \
    && mv /go/bin/godotenv /app/bin

WORKDIR /app
