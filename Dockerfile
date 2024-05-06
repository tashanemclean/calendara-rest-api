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
ENV PORT="9000"
ENV ENVIRONMENT="staging"
ENV API_BASE_URL="http://api.backend.calendara.io:5000"
ENV DATABASE_CONNECTION_URL="postgres://postgres:postgres@postgres_db:5432/calendara?ssslmode=disables&search_path=calendara"
ARG COMMIT_REF
ENV COMMIT_REF $COMMIT_REF

# Copy over necessary application files
COPY --from=builder ./app/bin /app/bin

# Copy db migrations directory
COPY --from=builder /app/internal/db/migrations /app/internal/db/migrations

EXPOSE 9000

ENTRYPOINT [ "/app/bin/calendara_rest_api" ]