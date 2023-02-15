# build stage
FROM golang:alpine AS builder

# RUN apk add build-base

WORKDIR /go/src/github.com/zalgonoise/eljoth-go-code-review/coupon-service
COPY ./ ./

RUN go mod download
RUN mkdir /build \
    && go build -o /build/coupon-service ./cmd/coupon_service \
    && chmod +x /build/coupon-service

# final stage
FROM alpine:edge
COPY --from=builder /build/coupon-service /coupon-service

ARG COUPON_SERVICE_PORT
ENV COUPON_SERVICE_PORT=${COUPON_SERVICE_PORT:-8080}

EXPOSE ${COUPON_SERVICE_PORT}
CMD ["/coupon-service"]
