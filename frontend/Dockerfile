FROM node:12.18.3-alpine3.12 as build

RUN apk update
RUN apk upgrade
RUN apk add --no-cache git

WORKDIR /build

COPY package.json .
COPY package-lock.json .
RUN npm ci

COPY . .
RUN npm run build

FROM nginx:1.19.3 as runner

COPY nginx.conf /etc/nginx/nginx.conf
COPY mime.types /etc/nginx/conf/mime.types
COPY public/ /var/www

COPY --from=build /build/dist /var/www
EXPOSE 80

CMD ["nginx"]
