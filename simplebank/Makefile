postgres:
	docker run --name postgres17_new -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=ARga12@@ -d postgres:17-alpine

createdb:
	docker exec -it postgres17_new createdb --username=root --owner=root simplebank

migrateup:
	migrate -path db/migration -database "postgresql://root:ARga12%40%40@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:ARga12%40%40@localhost:5432/simplebank?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgres17_new dropdb --username=root simplebank

sqlc:
	sqlc generate

.PHONY: createdb dropdb postgres migrateup migratedown sqlc