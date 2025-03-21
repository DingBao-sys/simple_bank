postgres12image:
	docker pull postgres:12-alpine

postgres12:
	docker run --name postgres12 -p 3456:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:3456/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:3456/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

mockgen:
	mockgen -package mockdb -destination ./db/mock/store.go github.com/DingBao-sys/simple_bank/db/sqlc Store

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/DingBao-sys/simple_bank/db/sqlc Store

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:3456/simple_bank?sslmode=disable" -verbose down 1

migrateup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:3456/simple_bank?sslmode=disable" -verbose up 1

createcontainer:
	docker build -t simple:latest .

# createmigration:
# 	migrate create -seq sql -dir db/migrate -seq 
.PHONY: postgres12 createdb dropdb migrateup migratedown sqlc test mockgen server mock migratedown1 migrateup1 createcontainer