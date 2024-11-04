# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod tidy
COPY ./ ./

RUN go build -o /wellbe-common

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /wellbe-common /wellbe-common
COPY --from=build /app/settings /settings
EXPOSE 8082

USER nonroot:nonroot

ENTRYPOINT ["/wellbe-common"]