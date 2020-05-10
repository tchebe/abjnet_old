compile_proto:
	protoc  --go_out=. --micro_out=. proto/product/product.proto
build: 
		docker build -t abjnet_product_service .
run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 abjnet_product_service