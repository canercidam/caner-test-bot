.PHONY: build
build:
	@CGO_ENABLED=0 go build -o build/bot .
	@docker build -t disco.forta.network/caner-test-bot -f Dockerfile ./build
	@docker push disco.forta.network/caner-test-bot

.PHONY: build-nop
build-nop:
	@CGO_ENABLED=0 go build -o build/bot ./nop
	@docker build -t disco.forta.network/caner-test-bot -f Dockerfile ./build
	@docker push disco.forta.network/caner-test-bot
