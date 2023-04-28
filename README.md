# Bank of Asia

one bank, all assets.

## Progress
* created Postgres DB schema
* added migration support using [`golang-migrate`](https://github.com/golang-migrate/migrate#cli-usage)
* created Makefile for docker container creation, db up/down migrations
* added tests for DB transactions, preventing deadlocks
* added Git workflow for tests for `main` branch

## DB
* [`sqlc`](https://sqlc.dev/) to translate SQL into Go code
* [`testify`](https://github.com/stretchr/testify) to ease test writing
* [`migrate`](https://github.com/golang-migrate/migrate) to migrate DB versions up/down
