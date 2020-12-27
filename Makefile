compose-app:
	- docker-compose -f docker-compose-app.yml up 

compose-dbs:
	- docker-compose -f docker-compose-dbs.yml up

gen:
	- @echo "generating..."
	- go generate ./graph/...
	- @echo "done generating! ✅"

start:
	- go run server.go

build:
	- go build -o main .


.PHONY: compose-app compose-dbs gen build