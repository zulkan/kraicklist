# Builder
FROM golang:1.16.4-alpine3.13 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 3001

COPY --from=builder /app/engine /app
COPY --from=builder /app/data.gz /app
COPY --from=builder /app/kraicklist/delivery/static /app/kraicklist/delivery/static

CMD /app/engine