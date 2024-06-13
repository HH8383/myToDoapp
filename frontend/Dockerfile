FROM node:latest AS build-react

WORKDIR /app/frontend

COPY package.json ./
RUN yarn install

COPY . .

RUN yarn build

EXPOSE 3000

CMD ["yarn","start"]