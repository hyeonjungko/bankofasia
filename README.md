# Bank of Asia

one bank, all assets.

## Progress
* created Postgres DB schema
* added migration support using [`golang-migrate`](https://github.com/golang-migrate/migrate#cli-usage)
* created Makefile for docker container creation, db up/down migrations
* added tests for DB transactions, preventing deadlocks
* added Git workflow for tests for `main` branch
* added server routes (+ tests) for `account` related functionality
* added server routes for `transfer` related functionality

## External Packages
* [`sqlc`](https://sqlc.dev/) to translate SQL into Go code
* [`testify`](https://github.com/stretchr/testify) to ease test writing
* [`migrate`](https://github.com/golang-migrate/migrate) to migrate DB versions up/down
* [`viper`](https://github.com/spf13/viper) to read/set env & configuration variables

## TODO Backlog
* Add tests for all errors of `/transfers` to gain full coverage
* Add tests for all errors of `/users` to gain full coverage
