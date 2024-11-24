FROM node:21 as builder
WORKDIR /var/www/famiphoto_frontend
COPY ../.. .
RUN npm ci
RUN npm run build

FROM node:21
WORKDIR /var/www/famiphoto_frontend
ADD .env ./
COPY --from=builder /var/www/famiphoto_frontend/.output ./
EXPOSE 3000
CMD ["node", "./server/index.mjs"]
