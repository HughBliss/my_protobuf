PROTO_DIR := proto
GEN_DIR   := ./gen

GO_MODULE_NAME := github.com/hughbliss/my_protobuf/gen
PROTO_FILES    := $(wildcard $(PROTO_DIR)/**/*.proto $(PROTO_DIR)/*.proto)

GO_OUT_OPTIONS           := module=$(GO_MODULE_NAME)
GO_GRPC_OUT_OPTIONS      := require_unimplemented_servers=false,module=$(GO_MODULE_NAME)
GRPC_GATEWAY_OUT_OPTIONS := logtostderr=true,module=$(GO_MODULE_NAME)
OPENAPI_OUT_OPTIONS      := logtostderr=true,allow_merge=true,output_format=yaml

SWAGGER_DIR := ./swagger

.PHONY: all
all: go-gen

.PHONY: go-gen
go-gen: swagger go-grpc go-grpc-gateway

.PHONY: swagger go-grpc go-grpc-gateway
swagger:
	mkdir -p $(SWAGGER_DIR)
	protoc -I $(PROTO_DIR) \
		--openapiv2_out=$(OPENAPI_OUT_OPTIONS):$(SWAGGER_DIR) \
		$(PROTO_FILES)

go-grpc:
	protoc -I $(PROTO_DIR) \
		--go_out=$(GO_OUT_OPTIONS):$(GEN_DIR)/ \
		$(PROTO_FILES)

go-grpc-gateway:
	protoc -I $(PROTO_DIR) \
		--go-grpc_out=$(GO_GRPC_OUT_OPTIONS):$(GEN_DIR)/ \
		--grpc-gateway_out=$(GRPC_GATEWAY_OUT_OPTIONS):$(GEN_DIR)/ \
		$(PROTO_FILES)
