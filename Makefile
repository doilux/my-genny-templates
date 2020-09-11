.PHONY: test
test:
	go generate && go test ./...

.PHONY: build
build:test
	go build
