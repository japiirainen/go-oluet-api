compose-app:
	- docker-compose -f docker-compose-app.yml up 

compose-dbs:
	- docker-compose -f docker-compose-dbs.yml up

gen:
	- @echo "generating..."
	- go generate ./graph/...
	- @echo "done generating! ✅"

start:
	- go run --race server.go

build:
	- go build --race -o main .

migrate-up:
	- migrate -database postgres://dev:dev@localhost:5432/oluet_api_dev?sslmode=disable -path db/migrations/postgres up

migrate-down:
	- migrate -database postgres://dev:dev@localhost:5432/oluet_api_dev?sslmode=disable -path db/migrations/postgres down


.PHONY: compose-app compose-dbs gen build
