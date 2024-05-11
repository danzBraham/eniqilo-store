build:
	@go build -o bin/eniqilo-store cmd/server/main.go

run: build
	@./bin/eniqilo-store