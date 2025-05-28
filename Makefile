.PHONY: run
run:
	@echo "Running go-monolith"
	@docker-compose -f docker-compose.yml up -d

.PHONY: restart
restart:
	docker-compose -f docker-compose.yml restart backend

.PHONY: down
down:
	docker-compose -f docker-compose.yml down

.PHONY: log
log:
	docker-compose -f docker-compose.yml logs --follow backend

.PHONY: test
test:
	@go test -failfast -race $(if $(JSON_OUTPUT),-json) ./...

.PHONY: generate
generate: generate-graphql

.PHONY: generate-graphql
generate-graphql: install-gql-tools
	cd ./api/graphql && gqlgen generate

.PHONY: install-gql-tools
install-gql-tools:
	GOFLAGS="" go install github.com/99designs/gqlgen@v0.17.72