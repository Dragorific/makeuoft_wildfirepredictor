# Alpine linux with timezone and CA data
FROM golang:alpine as goalpine

RUN adduser -D -g '' gopher

# certificates + timezone data
RUN apk update
RUN apk --no-cache add ca-certificates tzdata git


# build image 
FROM goalpine as build
ARG block

COPY . /oom
WORKDIR /oom
RUN GIT_COMMIT=$(git rev-parse --short HEAD) && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -mod=vendor -installsuffix cgo -ldflags="-w -s -X main.versionHash=$GIT_COMMIT" -o /oom/entrypoint ${block}/main.go


# final image
FROM alpine

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /oom /
USER gopher

ENTRYPOINT ["/entrypoint"]