.PHONY: run
run: cmd/server/main.go
	go build -o dist/server/main ./cmd/server/main.go
	./dist/server/main

.PHONY: test
test:
	go test