FROM node:0.12.7

MAINTAINER Adam Hooper <adam@adamhooper.com>

RUN groupadd -r app && useradd -r -g app app

COPY . /opt/app/

USER app
WORKDIR /opt/app

ENV PORT 3000
EXPOSE 3000
CMD [ "node", "server.js" ]
