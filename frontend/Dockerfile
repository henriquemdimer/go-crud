FROM node AS prod

WORKDIR /app/client

COPY package*.json ./
RUN npm install
COPY . .

RUN npm run build

FROM nginx:alpine AS web

WORKDIR /usr/share/nginx/html

COPY --from=prod /app/client/nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=prod /app/client/build .

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]