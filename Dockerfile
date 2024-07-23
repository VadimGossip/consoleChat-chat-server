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
COPY --from=build /build/config/config.yml /app/config/config.yml

CMD ["./chat-server"]