# почему то работает не так как запланировано, compose поднимается но

FROM node:lts-alpine as develop-stage

WORKDIR /app

COPY package*.json ./

RUN npm i -g @quasar/cli@latest
COPY . .
RUN yarn
#RUN quasar dev

##############################################################################
## develop stage
#FROM node:lts-alpine as develop-stage
#WORKDIR /app
#COPY package*.json ./
#RUN yarn global add @quasar/cli
#COPY . .
#////////////////////////////
## build stage
#FROM develop-stage as build-stage
#RUN yarn
#RUN quasar build \
#
## production stage
#FROM nginx as production-stage
#COPY --from=build-stage /app/dist/spa /usr/share/nginx/html
#EXPOSE 80
#CMD ["nginx", "-g", "daemon off;"]


































#CMD ["quasar", "dev"]

#  develop stage
#FROM node:lts-alpine as develop-stage
#WORKDIR /src
#COPY package*.json ./
#CMD npm i -g @quasar/cli@latest
#COPY . .
#RUN quasar dev

##local-deps
#FROM develop-stage
#COPY . .
#RUN quasar dev
#    RUN yarn

# # build stage (spa)
# FROM local-deps-stage as build-stage-spa
# RUN quasar build -m spa

# # production stage (spa)
# FROM nginx:stable-alpine as production-stage-spa
# COPY --from=build-stage-spa /src/dist/spa /usr/share/nginx/html
# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]

# # build stage (ssr)
# FROM local-deps-stage as build-stage-ssr
# RUN quasar build -m ssr

# #production stage (ssr)
# FROM node:lts-alpine as production-stage-ssr
# WORKDIR /app
# COPY --from=build-stage-ssr /src/dist/ssr .
# RUN yarn
# EXPOSE 3000
# CMD ["node", "index.js"]
