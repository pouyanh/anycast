FROM golang:latest as builder

ARG PROJECT_NAME=servant
ARG PROJECT_NAMESPACE=github.com/pouyanh/anycast

WORKDIR $GOPATH/src/$PROJECT_NAMESPACE/$PROJECT_NAME
COPY . .
RUN go generate ./...
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
			go build -a -ldflags '-w -s' -o /build/$PROJECT_NAME

FROM alpine:latest
MAINTAINER Pouyan Heyratpour <pouyan@janstun.com>

ARG PROJECT_NAME=servant

# Copy compiled binaries and modules
COPY --from=builder /build/$PROJECT_NAME /app/main

ENTRYPOINT ["/app/main"]
