# Bank of Asia

one bank, all assets.

## Progress
* created Postgres DB schema
* added migration support using [`golang-migrate`](https://github.com/golang-migrate/migrate#cli-usage)
* created Makefile for docker container creation, db up/down migrations

## DB
* using [`sqlc`](https://sqlc.dev/) to translate SQL into Go code

## Wanted Features
* holds all currencies
* holds all crypto assets
* online and offline transactions


