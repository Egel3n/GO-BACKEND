postgres:
	docker run --name myPost -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345 -d postgres:16.1-bullseye

createdb:
	docker exec -it myPost createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it myPost dropdb simple_bank

migrateup:
	 migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc