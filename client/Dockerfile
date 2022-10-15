# syntax=docker/dockerfile:1

FROM node:18.11

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 8100

CMD [ "npm", "start" ]
