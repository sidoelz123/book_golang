FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
COPY .env .env
RUN go mod tidy
RUN go build -o binary

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/binary /app/
ENTRYPOINT ["/app/binary"]

