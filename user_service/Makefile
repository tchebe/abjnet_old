compile_proto:
	protoc  --go_out=. --micro_out=. proto/user/user.proto
build: 
		docker build -t abjnet_user_service .
run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 abjnet_user_service