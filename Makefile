DB_URL=postgresql://root:secret@localhost:5432/stickverse?sslmode=disable

network:
	docker network create stickverse-network

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

mysql:
	docker run --name mysql8 -p 3306:3306  -e MYSQL_ROOT_PASSWORD=secret -d mysql:8

createdb:
	docker exec -it postgres createdb --username=root --owner=root stickverse

dumpdb:
	docker exec -it postgres pg_dump --username=root stickverse
dropdb:
	docker exec -it postgres dropdb stickverse

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

proto:
	protoc -I ./docs/proto --go_out=. ./docs/proto/common.proto
	protoc -I ./docs/proto --go_out=. --go-grpc_out=. ./docs/proto/auth.proto
	protoc -I ./docs/proto --go_out=. --go-grpc_out=. ./docs/proto/user.proto
	protoc -I ./docs/proto --go_out=. --go-grpc_out=. ./docs/proto/maplist.proto
	protoc -I ./docs/proto --go_out=. --go-grpc_out=. ./docs/proto/structure.proto


evans:
	evans --host localhost --port 9090 -r repl

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 db_docs db_schema sqlc test server mock proto evans
