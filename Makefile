postgres:
	docker run --name myPost -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345 -d postgres:16.1-bullseye

createdb:
	docker exec -it myPost createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it myPost dropdb simple_bank

migrateup:
	 migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	 migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	 migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:12345@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store	

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock migrateup1 migratedown1