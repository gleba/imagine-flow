FROM golang:1.18beta1-alpine3.14 as builder

RUN apk add musl-dev vips-dev gcc

WORKDIR /src
COPY . .
ARG APP_VER=local
RUN go build -ldflags "-X imagine-flow/vars.Version=$APP_VER" -o /bin/app /src/main.go
#
FROM alpine:3.14
RUN apk add vips-dev
COPY --from=builder /bin/app /app
CMD /app