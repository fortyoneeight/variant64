# syntax=docker/dockerfile:1

FROM node:18.11-alpine3.16

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8001

CMD [ "npm", "start" ]
