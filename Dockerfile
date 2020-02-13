FROM alpine:latest

RUN apk add bash curl

COPY install.sh /app/install.sh

WORKDIR /app

RUN chmod +x /app/install.sh
RUN /app/install.sh

CMD ["gws"]
