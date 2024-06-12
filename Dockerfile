FROM node:latest AS build-react
WORKDIR /app/frontend
COPY frontend/package.json ./
RUN yarn
COPY frontend ./
RUN yarn build

FROM golang:latest AS build-golang
WORKDIR /app/backend
COPY backend .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build-react /app/frontend/build ./frontend
COPY --from=build-golang /app/backend/main ./backend
EXPOSE 8080
CMD ["./backend"]