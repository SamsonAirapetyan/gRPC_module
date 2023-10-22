.PHONY: protos

protos:
	protoc -I protos/ protos/currency.proto --go-grpc_out=. protos/currency

protoss:
	protoc --go_out=. --go_opt=paths=import --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=import protos/currency.proto