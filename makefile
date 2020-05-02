.PHONY: build
build:
	go build -v
	./crm -config=config.json
.DEFAULT_GOAL := build