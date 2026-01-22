# STEP-1
# build app from source

FROM golang:1.25.5-alpine3.21 AS builder

WORKDIR /mysource

COPY ./go.mod ./go.sum ./
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./docs ./docs
COPY ./vendor ./vendor

RUN go build -o app ./cmd/main.go

# STEP-2
# make container

FROM alpine:3.21
RUN apk add --no-cache curl

RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

USER appuser

WORKDIR /myapp

COPY --from=builder /mysource ./

CMD [ "/myapp/app" ]
