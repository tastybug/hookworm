check-fmt:
	docker run --rm -v ${PWD}:/app -w /app golang:1.24-alpine sh -c 'gofmt -l . | grep . && exit 1 || exit 0'
fmt:
	go fmt ./...
lint:
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run ./...
build:
	go build cmd/cmdline/main.go && mv main hookworm
install:
	sudo mv ./hookworm /usr/local/bin/
