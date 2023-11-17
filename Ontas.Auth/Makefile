proto:
	protoc -I=./ --go-grpc_out=require_unimplemented_servers=false:.  --go_out=.  pkg/pb/*.proto
server:
	go run cmd/main.go