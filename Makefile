include .env

.PHONY: run # Run the Last-Order API
run:
	docker compose -f compose.yaml build --no-cache
	docker compose -f compose.yaml up -d

.PHONY: stop # Stop the Last-Order API
stop:
	docker compose -f compose.yaml down

.PHONY: build-development # Build the Last-Order API in development mode
build-development:
	docker compose -f compose.development.yaml build --no-cache

.PHONY: start-development # Start the Last-Order API in development mode
start-development:
	docker compose -f compose.development.yaml up

.PHONY: run-development # Build&Start the Last-Order API in development mode
run-development:
	docker compose -f compose.development.yaml build --no-cache
	docker compose -f compose.development.yaml up

.PHONY: stop-development # Stop the Last-Order API in development mode
stop-development:
	docker compose -f compose.development.yaml down --rmi all --volumes --remove-orphans

.PHONY: lint # Run the Linter
lint:
	golangci-lint run
	docker compose -f compose.development.yaml exec lint 

.PHONY: test # Run the tests
test:
	go test -v ./...

.PHONY: coverage # Run the tests with coverage
coverage:
	go test -v -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
	open cover.html

.PHONY: migrate-up-development
migrate-up-development:
	docker build -t rnp/migrate -f docker/development/migrate/Dockerfile .
	docker run --rm --net=last-order_development-lo-network \
	-v $$PWD:/work -w /work rnp/migrate \
	-path pkg/database/migration \
	-database 'postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable' --verbose up
	
.PHONY: migrate-down-development
migrate-down-development:
	docker build -t rnp/migrate -f docker/development/migrate/Dockerfile .
	docker run --rm --net=last-order_development-lo-network \
	-v $$PWD:/work -w /work sql/migrate \
	-path pkg/database/migration \
	-database 'postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable' --verbose down

.PHONY: sqlboiler
sqlboiler:
	docker build -t rnp/sqlboiler -f docker/development/sqlboiler/Dockerfile .
	docker run --rm --net=last-order_development-lo-network \
	-v $$PWD:/work -w /work rnp/sqlboiler \
	sqlboiler psql -c internal/database/sqlboiler.toml -o pkg/database/model

.PHONY: gen-db-doc
gen-db-doc:
	docker build -t rnp/tbls -f docker/development/tbls/Dockerfile .
	docker run --rm --net=$(DOCKER_NETWORK) -v $$PWD:/work -w /work rnp/tbls -c .tbls.yml doc $(DB_URL) --rm-dist

.PHONY: clean
clean:
	docker prune -f
