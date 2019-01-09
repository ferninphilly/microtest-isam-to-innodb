# Build stage
FROM golang:alpine AS build-env

COPY . /go/src/goinnodb
WORKDIR /go/src/goinnodb
RUN apk update && \
    apk upgrade

# Install dep if needed
# ENV DEP_VERSION="0.4.1"
# RUN curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $GOPATH/bin/dep
# RUN chmod +x $GOPATH/bin/dep
# RUN dep ensure

# # Build your app
# RUN go build -o goinnodb

# # Final stage
# FROM alpine

# WORKDIR /app/goinnodb
# COPY --from=build-env /go/src/your/project/path /app/goinnodb
# ENTRYPOINT ["/app/myapp/myapp"]
