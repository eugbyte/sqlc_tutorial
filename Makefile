.PHONY: run

run:
	go run ./cmd/main.go

sqlc:
	sqlc generate

create-migration:
	goose -dir=internal/db/migration -s create sql_filename sql

apply-migration:
	goose -dir=internal/db/migration postgres "user=postgres password=postgres dbname=sqlc_tutorial sslmode=disable" up

rollback-migration-to:
	goose -dir=internal/db/migration postgres "user=postgres password=postgres dbname=sqlc_tutorial sslmode=disable" down-to 123456