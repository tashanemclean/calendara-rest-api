# Stage 1) Install dependencies and build application
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache make upx git

ARG GITHUB_ID
ARG GITHUB_TOKEN
ARG DOCKER_TAG
ENV BUILD_TAG=$DOCKER_TAG

# copy over all source code to be built
COPY . /app
RUN make go-install
RUN make go-build
RUN make go-upx

# Stage 2) package and run application
FROM alpine:latest

ARG DOCKER_TAG
ENV BUILD_TAG=$DOCKER_TAG
ARG COMMIT_REF
ENV COMMIT_REF $COMMIT_REF

# Copy over necessary application files
COPY --from=builder ./app/bin /app/bin

# Copy db migrations directory
COPY --from=builder /app/internal/db/migrations /app/internal/db/migrations

# expose port 80
EXPOSE 9000

ENTRYPOINT [ "/app/bin/calendara_rest_api" ]