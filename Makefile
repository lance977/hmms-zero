
GOCTL=goctl

SVC=user # Service

gen_proto:
	$(GOCTL) rpc protoc ./rpc/$(SVC)/$(SVC).proto --go_out=./rpc/$(SVC) --go-grpc_out=./rpc/$(SVC) --zrpc_out=./rpc/$(SVC)
