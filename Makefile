build:
	@go build -o ./bin/rest

run: build
	./bin/rest

test:
	@go test -v ./...
