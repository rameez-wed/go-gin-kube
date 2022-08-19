postgres:
	docker run --name postgres-author --network author-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14.4-alpine

createdb:
	docker exec -it postgres-author createdb --username=root --owner=root author

dropdb:
	docker exec -it postgres-author dropdb author

migrateup:
	migrate -database "postgres://root:password@localhost:5432/author?sslmode=disable" -path db/internal/migrations -verbose up

migratedown:
	migrate -database "postgres://root:password@localhost:5432/author?sslmode=disable" -path db/internal/migrations -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/go-gin-kube/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock