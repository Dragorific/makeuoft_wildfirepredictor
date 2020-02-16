# base image
FROM node:12.2.0-alpine

# set working directory
WORKDIR /frontend

# install and cache web dependencies
RUN npm config set unsafe-perm true
RUN npm install
COPY package.json /frontend/package.json
RUN npm install @vue/cli@3.7.0 -g

# start web
CMD ["npm", "run", "build"]