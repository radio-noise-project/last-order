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
	docker compose up db -d
	docker compose exec db /bin/psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}

DB_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db/${POSTGRES_DB}?sslmode=disable"
DOCKER_NETWORK="last-order_db-network"
.PHONY: migrate-create
migrate-create: # make migrate-create name=sql_file_name
	migrate create -ext sql -dir migration -seq $(name) || echo "Please run: make migrate-create name=sql_file_name"

.PHONY: migrate-up
migrate-up:
	docker run --rm --rm --net=${DOCKER_NETWORK} -v ./migration:/migration migrate/migrate -path=/migration/ -database $(DB_URL) up $(limit) 		

.PHONY: migrate-down
migrate-down:
	docker run --rm --net=${DOCKER_NETWORK} -v ./migration:/migration -it migrate/migrate -path=/migration/ -database $(DB_URL) down $(limit)

.PHONY: migrate
migrate: #make migrate args="arg"
	docker run --rm --net=$(DOCKER_NETWORK) -v ./migration:/migration -it migrate/migrate -path=/migration/ -database $(DB_URL) $(arg) 

.PHONY: gen-db-doc
gen-db-doc:
	docker run --rm --net=$(DOCKER_NETWORK) -v $$PWD:/work -w /work ghcr.io/k1low/tbls -c .tbls.yml doc $(DB_URL) --rm-dist

.PHONY: clean
clean:
	rm cover.html cover.out 


