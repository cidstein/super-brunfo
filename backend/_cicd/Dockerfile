FROM alpine:3.17.0
LABEL maintainer="Alcides Souza <alcides.braga@gmail.com>"

RUN apk add tzdata --no-cache && \
	cp /usr/share/zoneinfo/UTC /etc/localtime && \
	printf "UTC\n" > /etc/timezone && \
	rm -rf /var/cache/apk/*

RUN addgroup \
	-g 1729 \
	app \
	&& adduser \
	-s /bin/sh \
	-D \
	-G app \
	-u 1729 \
	-H \
	app

COPY dist/super-brunfo /app/bin/super-brunfo

RUN find /app | xargs chgrp app
RUN find /app -type f | xargs chmod 0640
RUN find /app -type d | xargs chmod 0750

RUN chmod 0750 /app/bin/super-brunfo

USER app

WORKDIR /app
