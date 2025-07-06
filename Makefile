build:
	@go build -o bin/ecom cmd/migrate/main.go

test:
	@go test -d ./...

run: build
	@./bin/ecom
