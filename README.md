# Bank of Asia

one bank, all assets.

## Progress

- Postgres DB schema
- Migration support using [`golang-migrate`](https://github.com/golang-migrate/migrate#cli-usage)
- Makefile for docker container creation, db up/down migrations
- Tests for DB transactions, preventing deadlocks
- Git workflow for tests for `main` branch
- Server routes (+ tests) for `account` related functionality
- Server routes (+ tests) for `transfer` related functionality
- JWT token authentication
- PASETO token authentication
- `users/login` for user login using PASETO token authentication
- PASETO token authentication for `/accounts`, `/transfer` routes
- Dockerfile (multi-stage to reduce image size)
- Docker Compose file to migrate and start DB before running main api
- Set Up AWS (RDS, Secrets Manager, IAM User/User Groups, etc.)
- CI workflow to deploy Docker image to AWS Container Registry using OpenID Connect

## External Packages

- [`sqlc`](https://sqlc.dev/) to translate SQL into Go code
- [`testify`](https://github.com/stretchr/testify) to ease test writing
- [`migrate`](https://github.com/golang-migrate/migrate) to migrate DB versions up/down
- [`viper`](https://github.com/spf13/viper) to read/set env & configuration variables
- [`o1egl/paseto`](https://github.com/o1egl/paseto) to add PASETO token authentication

## TODO Backlog

- Add tests for all errors of `/transfers` to gain full coverage
- Add tests for all errors of `/users` to gain full coverage
