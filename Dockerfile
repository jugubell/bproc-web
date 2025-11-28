# BUILD SVELTE FRONTEND
FROM node:20-alpine AS fe_builder

ARG GUI_REPO=https://github.com/jugubell/bproc-web

RUN apk add --no-cache git

WORKDIR /app/bproc-web
RUN git clone ${GUI_REPO} .
WORKDIR /app/bproc-web/web
RUN npm install
RUN npm run build

# BUILD GO BACKEND
FROM golang:1.25-alpine AS be_builder

ARG GUI_REPO=https://github.com/jugubell/bproc-web

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk add --no-cache git

WORKDIR /app/bproc-web
RUN git clone ${GUI_REPO} .
RUN go mod download
RUN go build -o /app/bprocserver

# CLI JAR BUILD
FROM maven:3.9.11-eclipse-temurin-17-alpine AS jar_builder

ARG CLI_REPO=https://github.com/jugubell/bproc-cli

RUN apk add --no-cache git

WORKDIR /app/bproc-cli
RUN git clone ${CLI_REPO} .
RUN mvn dependency:go-offline
RUN mvn package -DskipTests

# PROD IMAGE
FROM alpine:3.22.1

RUN apk add openjdk17-jre-headless

WORKDIR /app

COPY --from=be_builder /app/bprocserver /usr/local/bin/bprocserver
COPY --from=fe_builder /app/bproc-web/web/dist/ /app/web/
COPY --from=jar_builder /app/bproc-cli/target/*.jar /app/libs/bproc.jar
COPY .env.bprocprod ./.env
COPY ./examples /app/

EXPOSE 8998

CMD ["/usr/local/bin/bprocserver"]