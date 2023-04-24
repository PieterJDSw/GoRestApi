build:
	@go build -o bin/gorestapi
run: build
	@./bin/gorestapi
test: 
	@go test -v ./...