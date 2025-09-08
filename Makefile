.DEFAULT_GOAL := test

# Run all tests
.PHONY: test
test:
	go test -v ./...

# Run tests with coverage
.PHONY: cover
cover:
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run fuzzy tests
.PHONY: fuzz
fuzz:
	@echo "Running fuzz tests..."
	go test -fuzz=FuzzWords$$ -fuzztime=10s
	go test -fuzz=FuzzWordsFromString$$ -fuzztime=10s
	go test -fuzz=FuzzConsistency$$ -fuzztime=10s
	go test -fuzz=FuzzPropertyBasedTesting$$ -fuzztime=10s

# Clean generated files
.PHONY: clean
clean:
	rm -f coverage.out coverage.html
	rm -rf testdata/fuzz/*/
