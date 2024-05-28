### Stage 1 - Build
###
### Build executable from code

FROM node:20-alpine as build
WORKDIR /app

### Copy files and generate build
COPY package.json .
COPY yarn.lock .
RUN yarn install
COPY . .
RUN yarn build

### Stage 2 - Install
###
### Build customize docker image for go app
FROM nginx:1.20.0

### Copy build
RUN rm -r /usr/share/nginx/html/*
COPY --from=build /app/dist/ /usr/share/nginx/html/

ENTRYPOINT ["nginx", "-g", "daemon off;"]