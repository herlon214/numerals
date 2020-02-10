run:
	@go run cmd/numerals/main.go

test:
	@go test -race -coverpkg=./pkg/... -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

bench:
	@go test -bench=. ./pkg/...