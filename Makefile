postgres:
	docker run --name postgres-author -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14.4-alpine

createdb:
	docker exec -it postgres-author createdb --username=root --owner=root author

dropdb:
	docker exec -it postgres-author dropdb author

migrateup:
	migrate -database "postgres://root:password@localhost:5432/author?sslmode=disable" -path internal/migrations -verbose up

migratedown:
	migrate -database "postgres://root:password@localhost:5432/author?sslmode=disable" -path internal/migrations -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc