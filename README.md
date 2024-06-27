# calendara-rest-api-api

## Section 1: Local Development Environment Setup

This microservice is designed as to be the REST API gateway for the service
[calendara-backend](https://github.com/tashanemclean/calendara-backend)

### Setup Local Environment

To use this miroservice, you need to have the following tools installed on your
developer machine:

1. go [Golang](https://go.dev/doc/install)

### Configure Environment variables

To use environment variables, create configs directory in project root, place
dev.json inside it.

```
configs/dev.json
```

```
"PORT": "80",
"ENVIRONMENT": "dev",
"API_BASE_URL": "http://localhost:9000"
```

### Running the app

To run , run:

```
$ go run main.go
```

### Running in docker

The port need to be explicitly set in `main.go` when testing in docker
environment

### Usage example

The ClassifyText api can be tested via http POST to /v1/process with text as
payload.

```
curl --location 'http://localhost:8080/v1/process' \
--header 'Content-Type: application/json' \
--data '{
    "text": "When was the movie Citizen Kane released?"
}'
```

## Limitations

Take care when exceeding your plan quota, the api will fail and you may need to
adjust your billing details,
[Read the docs](https://platform.openai.com/docs/guides/error-codes/api-errors)

## Notes

Environment variables passed into containers are strictly for testing, it is
recommended that these are omitted in lieu of secrets.
[Read the docs](https://docs.docker.com/compose/environment-variables/set-environment-variables/)
