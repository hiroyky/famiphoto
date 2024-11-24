FROM node:22 as local
WORKDIR /var/www/famiphoto_frontend
COPY ../../web/ .
RUN npm i
EXPOSE 3000

FROM node:22 as builder
WORKDIR /var/www/famiphoto_frontend
COPY ../../web/ .
RUN npm ci
RUN npm run build

FROM node:22
WORKDIR /var/www/famiphoto_frontend
ADD .env ./
COPY --from=builder /var/www/famiphoto_frontend/.output ./
EXPOSE 3000
CMD ["node", "./server/index.mjs"]
