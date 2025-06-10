PROTO_DIR  := proto
GO_GEN_DIR := ./go/pkg

GO_MODULE_NAME := github.com/hughbliss/my_protobuf/go/pkg
PROTO_FILES    := $(wildcard $(PROTO_DIR)/**/*.proto $(PROTO_DIR)/*.proto)

GO_OUT_OPTIONS           := module=$(GO_MODULE_NAME)
GO_GRPC_OUT_OPTIONS      := require_unimplemented_servers=false,module=$(GO_MODULE_NAME)
GRPC_GATEWAY_OUT_OPTIONS := logtostderr=true,module=$(GO_MODULE_NAME)
OPENAPI_OUT_OPTIONS      := logtostderr=true,allow_merge=true,output_format=yaml

SWAGGER_DIR := ./swagger

.PHONY: all
all: gen-tools go-gen

.PHONY: go-gen
go-gen: swagger go-grpc go-grpc-gateway acman

.PHONY: swagger go-grpc go-grpc-gateway acman
swagger:
	mkdir -p $(SWAGGER_DIR)
	protoc -I $(PROTO_DIR) \
		--openapiv2_out=$(OPENAPI_OUT_OPTIONS):$(SWAGGER_DIR) \
		$(PROTO_FILES)

go-grpc:
	protoc -I $(PROTO_DIR) \
		--go_out=$(GO_OUT_OPTIONS):$(GO_GEN_DIR)/ \
		$(PROTO_FILES)

go-grpc-gateway:
	protoc -I $(PROTO_DIR) \
		--go-grpc_out=$(GO_GRPC_OUT_OPTIONS):$(GO_GEN_DIR)/ \
		--grpc-gateway_out=$(GRPC_GATEWAY_OUT_OPTIONS):$(GO_GEN_DIR)/ \
		$(PROTO_FILES)

acman:
	protoc -I $(PROTO_DIR) \
		--acman_out=$(GO_GEN_DIR)/gen/acman \
		$(PROTO_FILES)

.PHONY: gen-tools
gen-tools: install-acman

.PHONY: install-acman
install-acman:
	go install ./go/pkg/acman/cmd/protoc-gen-acman
