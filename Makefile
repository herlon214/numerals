run:
	@go run cmd/numerals/main.go

test:
	@go test -short -coverpkg=./pkg/... -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out