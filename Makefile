clean:
	@go clean -testcache

test:
	go test ./...

test-cover: clean
	go test -coverprofile=./coverage.out  ./...
	go tool cover -html=coverage.out -o coverage.html
