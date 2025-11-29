# syntax=docker/dockerfile:1.6

## =============== Build Stage ===============
FROM node:20-alpine AS build
WORKDIR /app
ENV CI=true

ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL=$VITE_API_BASE_URL

COPY frontend_vue/package*.json ./
RUN npm ci --no-audit --no-fund

COPY frontend_vue/. ./
RUN npm run build

## =============== Runtime Stage ===============
FROM nginx:1.27-alpine
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=build /app/dist ./
COPY frontend_vue/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s CMD wget -qO- http://localhost:5173 || exit 1
