postgres:
	docker run --name postgres-15 -p5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres-15 createdb --username=root --owner=root bankofasia

dropdb:
	docker exec -it postgres-15 dropdb bankofasia

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankofasia?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankofasia?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown