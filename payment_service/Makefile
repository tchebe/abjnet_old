compile_proto:
	protoc  --go_out=. --micro_out=. proto/payment/payment.proto
build: 
		docker build -t abjnet_payment_service .
run:
	docker run -p 50054:50051 -e MICRO_SERVER_ADDRESS=:50051 abjnet_payment_service