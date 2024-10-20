
build:
	@go build -o bin/ks

run: build
	./bin/ks

release:
	CGO_ENABLED=0 go build -ldflags "-s -w" -o /release/ks

test-coverage:
	@echo "running against `git version`"; \
	echo "" > $(COVERAGE_REPORT); \
	$(GOTEST) -coverprofile=$(COVERAGE_REPORT) -coverpkg=./... -covermode=$(COVERAGE_MODE) ./...
