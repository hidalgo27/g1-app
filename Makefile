.PHONY: run-dev run-test run-main test

run-dev:
	@GO_ENV=dev go run ./cmd/main.go

run-test:
	@GO_ENV=test go run ./cmd/main.go
#	@GO_ENV=test go run ./cmd/scripts/create_user.go

run-main:
	@GO_ENV=main go run ./cmd/main.go

test:
	@GO_ENV=test go test ./...