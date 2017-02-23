build:
	go build github.com/lukewendling/watchman
	cd consumer; go build

test:
	go test -v

format:
	gofmt -w .
