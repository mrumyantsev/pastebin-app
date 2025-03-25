ARG GOLANG_VERSION
ARG ALPINE_VERSION

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} as builder

ARG APP_NAME

WORKDIR /project

COPY . .

RUN go build -o ./server ./cmd/${APP_NAME}

FROM alpine:${ALPINE_VERSION}

WORKDIR /project

RUN apk update && \
    apk add ca-certificates

COPY --from=builder /project/server .

ENTRYPOINT [ "./server" ]
