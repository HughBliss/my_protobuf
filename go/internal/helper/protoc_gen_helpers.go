package helper

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
	"io"
	"os"
)

func DecodeRequest() *pluginpb.CodeGeneratorRequest {
	req := &pluginpb.CodeGeneratorRequest{}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	if err := proto.Unmarshal(data, req); err != nil {
		panic(err)
	}
	return req
}

func EncodeResponse(resp proto.Message) {
	data, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stdout.Write(data); err != nil {
		panic(err)
	}
}

func GetOption[T any](options *descriptorpb.MethodOptions, info *protoimpl.ExtensionInfo) *T {
	if !proto.HasExtension(options, info) {
		return nil
	}
	ext := proto.GetExtension(options, info)
	if option, ok := ext.(*T); ok {
		return option
	}
	return nil
}
