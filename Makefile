#!/usr/bin/make -f
containerProtoVer=0.13.0
containerProtoImage=ghcr.io/cosmos/proto-builder:$(containerProtoVer)
DOCKER := $(shell which docker)
proto-gen:
	@echo "Generating Protobuf files"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./scripts/protocgen.sh; 
