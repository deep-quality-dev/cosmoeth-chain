#!/usr/bin/make -f

BUILDDIR ?= $(CURDIR)/build
containerProtoVer=0.13.0
containerProtoImage=ghcr.io/cosmos/proto-builder:$(containerProtoVer)
DOCKER := $(shell which docker)

build-feeder:
	mkdir -p $(BUILDDIR)/
	go build -mod=readonly $(BUILD_FLAGS) -trimpath -o $(BUILDDIR) ./feeder/feeder.go;

proto-gen:
	@echo "Generating Protobuf files"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./scripts/protocgen.sh; 

