FROM alpine:latest
RUN apk add --no-cache go nodejs npm
RUN mkdir /bubuKabum
ADD bot /bubuKabum/bot
ADD converter /bubuKabum/converter
COPY .env /bubuKabum/bot/.env
WORKDIR /bubuKabum/converter
RUN go mod download
WORKDIR /bubuKabum/bot
RUN npm install
COPY entrypoint.sh /bin/entrypoint.sh
RUN chmod +x /bin/entrypoint.sh
ENTRYPOINT ["/bin/sh", "/bin/entrypoint.sh"]
