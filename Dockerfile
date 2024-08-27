##
## Stage 1
##
FROM golang:1.22-alpine AS build

WORKDIR /build
ADD . /build
RUN go build -o app cmd/main.go

##
## Stage 2
##
FROM alpine:latest

WORKDIR /app
COPY --from=build /build/app /app/chat-server

EXPOSE $APP_PORT
CMD ["./chat-server"]