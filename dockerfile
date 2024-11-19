FROM golang:1.23.3-alpine3.20 AS build-env
ADD . /src
RUN apk add make
RUN cd /src && make build


FROM alpine:3.20
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app