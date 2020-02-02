# base image
FROM golang:1.7-alpine

# adds to the environment path (probably)
ADD . /backend

# set working directory
WORKDIR /backend

# start web
CMD ["go", "run", "main.go"]