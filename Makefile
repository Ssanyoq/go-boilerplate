.PHONY: default
default:
	go run .

env:
	touch .env
	cp .default_env .env