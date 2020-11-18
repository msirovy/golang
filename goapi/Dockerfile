FROM debian:stable-slim

LABEL authors='msirovy@gmail.com'

COPY goapi-linux /usr/local/bin/goapi
COPY public /usr/share/public

ENTRYPOINT ["/usr/local/bin/goapi"]
