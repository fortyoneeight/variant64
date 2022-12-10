# syntax=docker/dockerfile:1

FROM node:18.11

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

EXPOSE 3000

CMD [ "npx", "serve", "build" ]
