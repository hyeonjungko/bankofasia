# Bank of Asia

one bank, all assets.

## Progress
* created Postgres DB schema
* added migration support using [`golang-migrate`](https://github.com/golang-migrate/migrate#cli-usage)
* created Makefile for docker container creation, db up/down migrations
* added tests for DB transactions

## DB
* [`sqlc`](https://sqlc.dev/) to translate SQL into Go code
* [`testify`](https://github.com/stretchr/testify) to ease test writing



