FROM node:22-alpine AS build

WORKDIR /app

COPY package*.json package-lock.json ./

RUN npm ci

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

ARG VITE_API_URL

ENV VITE_API_URL=${VITE_API_URL}

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]