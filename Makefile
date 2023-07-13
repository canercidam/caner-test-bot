.PHONY: build
build:
	@CGO_ENABLED=0 go build -o bot .
	@docker build -t disco.forta.network/caner-test-bot .
	@docker push disco.forta.network/caner-test-bot
