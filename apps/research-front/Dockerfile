FROM node:11-alpine as build
WORKDIR /app

# .dockerignoreしていないものは全てcopy
COPY ./ /app/

RUN yarn install \
    && yarn run build

FROM nginx:mainline-alpine
COPY --from=build /app/dist /app/
