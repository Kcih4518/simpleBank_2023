DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network

postgres:
					docker stop postgres
					docker rm postgres
					docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret --network bank-network -d postgres:14-alpine

createdb:
					docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
					docker exec -it postgres dropdb simple_bank

migrateup:
					migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
					migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
					migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
					migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
					sqlc generate

test:
					go test -count=1 -v -cover ./...

clean:
					docker exec -it postgres psql -U root -d simple_bank -c "TRUNCATE TABLE accounts, entries, transfers CASCADE;"

server:
					go run main.go

air:
					air -c .air.toml

mock:
					mockgen -build_flags=--mod=mod -destination db/mock/store.go -package mockdb github.com/Kcih4518/simpleBank_2023/db/sqlc Store

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc test clean server air mock migrateup1 migratedown1
