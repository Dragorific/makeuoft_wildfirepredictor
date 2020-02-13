FROM golang:latest AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
ARG block
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/entrypoint ${block}/main.go

FROM alpine:latest AS production
COPY --from=builder /app /
ENTRYPOINT ["/entrypoint"]