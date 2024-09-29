FROM golang:alpine AS base

WORKDIR /app
RUN mkdir /out

COPY go.mod .
COPY go.sum .
RUN go mod download


FROM base as build
ARG APP_NAME

WORKDIR /app
ADD . /app

RUN go build \
    -o /out/${APP_NAME} cmd/${APP_NAME}/main.go

FROM alpine as relise
WORKDIR /app
COPY --from=build /out/${APP_NAME} /app/

CMD ["/app/core-service"]