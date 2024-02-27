run: build  # Make 'run' depend on 'build'
	@./bin/api

build:
	@go build -o bin/api

test:
	@go test -v ./...
