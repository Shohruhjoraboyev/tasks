containerrun:
	docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=myPass -d postgres:12-alpine
migrateup:
	migrate -path db/migration -database "postgresql://postgres:myPass@localhost:5432/bankdb?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:myPass@localhost:5432/bankdb?sslmode=disable" -verbose down