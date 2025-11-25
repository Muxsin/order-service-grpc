run-orders:
	@go run services/orders/http.go services/orders/grpc.go services/orders/main.go

run-kitchen:
	@go run services/kitchen/http.go services/kitchen/main.go

gen:
	@protoc \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders --go-grpc_opt=paths=source_relative \
		protobuf/orders.proto
