FROM golang:alpine as builder

RUN apk update && apk upgrade && apk add --no-cache git

RUN go version

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

ARG buildtime_variable=default_value
ENV DB_HOST=$buildtime_variable

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consignment-service

#Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

ADD vars.env /app/vars.env


COPY --from=builder /app/consignment-service .

CMD ["./consignment-service"]