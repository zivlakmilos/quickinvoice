.PHONY: build

run:
	@air

build-dev:
	@GOOS=linux go build -o build/quickinvoice ./cmd/quickinvoice

build:
	@GOOS=linux go build -tags prod -o build/quickinvoice ./cmd/quickinvoice
