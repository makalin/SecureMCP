.PHONY: build clean test run docker-build docker-run

# Build the application
build:
	go build -o bin/securemcp cmd/securemcp/main.go

# Clean build artifacts
clean:
	rm -rf bin/
	rm -rf dist/
	rm -rf reports/

# Run tests
test:
	go test -v ./...

# Run the application
run: build
	./bin/securemcp

# Build Docker image
docker-build:
	docker build -t makalin/securemcp .

# Run Docker container
docker-run:
	docker run -p 8080:8080 makalin/securemcp

# Install dependencies
deps:
	go mod download
	go mod tidy

# Generate documentation
docs:
	godoc -http=:6060

# Lint code
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./... 