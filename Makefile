.PHONY: sql
sql:
	sqlc generate

.PHONY: swagger
swagger:
	@echo "Generate swagger documentation"
	swag init -g cmd/main.go --parseDependency --parseInternal

.PHONY: test
test:
	go test -v -race -coverprofile=coverage.out ./...
