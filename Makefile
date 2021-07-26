.PHONY: build
build: clean
	@go build -o build/cleanexample cmd/main.go

.PHONY: run
run: build
	@docker-compose up -d
	@./build/cleanexample

.PHONY: clean
clean:
	@docker-compose down --remove-orphans
	@rm -rf build/
