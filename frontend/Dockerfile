# base image
FROM node:12.2.0-alpine

# set working directory
WORKDIR /frontend

# add `/web/node_modules/.bin` to $PATH
ENV PATH /frontend/node_modules/.bin:$PATH

# install and cache web dependencies
COPY package.json /frontend/package.json
RUN npm install
RUN npm install @vue/cli@3.7.0 -g
RUN npm config set unsafe-perm true

# start web
CMD ["npm", "run", "build"]