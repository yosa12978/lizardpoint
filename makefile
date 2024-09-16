.DEFAULT_GOAL = build

build:
	@go mod tidy
	@go build -o bin/lizardpoint ./cmd/lizardpoint/main.go
	@cp config.yaml bin

run: build
	@LP_ADDR=0.0.0.0:5000 ./bin/lizardpoint