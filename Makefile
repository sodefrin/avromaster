PACKAGES = $(shell go list ./...)

test:
	@go test -v -parallel=4 $(PACKAGES)

lint:
	@golint $(PACKAGES)

vet:
	@go vet $(PACKAGES)

coverage:
	@go test -v -race -cover -covermode=atomic -coverprofile=coverage.txt $(PACKAGES)

.PHONY: test lint vet coverage
