.PHONY: sql
sql:
	sqlc generate

.PHONY: swagger
swagger:
	@echo "Generate swagger documentation"
	swag init -g cmd/main.go --parseDependency --parseInternal

.PHONY: test-integration
test-integration:
	cd ./tests/integration && go test -v
