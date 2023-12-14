FROM node:18 as build

WORKDIR /frontend
COPY package*.json ./
RUN npm i
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /frontend/build /usr/share/nginx/html
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]