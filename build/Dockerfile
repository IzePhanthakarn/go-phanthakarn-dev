FROM golang:1.18.0-alpine3.15 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates git
  

RUN mkdir /api
WORKDIR /api
ADD . /api
RUN go install github.com/swaggo/swag/cmd/swag@v1.8.0
RUN swag init

RUN go mod download
RUN go build -o api

FROM alpine:3.12.0

RUN apk update && apk upgrade && \
  apk add --no-cache ca-certificates tzdata

COPY --from=builder /api/api .
COPY --from=builder /api/docs /docs

ADD /configs /configs

EXPOSE 8000

ENTRYPOINT ["/api"]