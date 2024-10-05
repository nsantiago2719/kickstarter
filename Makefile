
build:
	@go build -o bin/ks

run: build
	./bin/ks

release:
	CGO_ENABLED=0 go build -ldflags "-s -w" -o /release/ks
