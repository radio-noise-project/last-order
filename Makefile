include .env

.PHONY: ci
ci:
	golangci-lint run

.PHONY: test
test:
	go test -v  ./...
.PHONY: coverage
coverage:
	go test -v -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
.PHONY: run
run:
	go run cmd/last-order/main.go 
.PHONY: psql
psql:
	docker compose up -d
	docker compose exec db /bin/psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}
DB_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db/${POSTGRES_DB}?sslmode=disable"

.PHONY: migrate-create
migrate-create: # make migrate-create name=sql_file_name
	docker run --rm --net=last-order_db-network -v ./migration:/migration migrate/migrate -path=/migration/ -database $(DB_URL)  create -ext sql -dir migration -seq $(name) 
.PHONY: migrate-up
migrate-up:
	docker run --rm --rm --net=last-order_db-network -v ./migration:/migration migrate/migrate -path=/migration/ -database $(DB_URL) up $(limit) 		


.PHONY: migrate-down
migrate-down:
	docker run --rm --net=last-order_db-network -v ./migration:/migration -it migrate/migrate -path=/migration/ -database $(DB_URL) down $(limit)

.PHONY: migrate-version
migrate: #make migrate args="arg"
	docker --rm run --net=last-order_db-network -v ./migration:/migration -it migrate/migrate -path=/migration/ -database $(DB_URL) $(arg) 

clean:
	rm cover.html cover.out 


