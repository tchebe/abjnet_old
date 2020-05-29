compile_proto:
	protoc  --go_out=. --micro_out=. proto/prestation/prestation.proto
build: 
		docker build -t abjnet_prestation_service .
run:
	docker run -p 50056:50051 -e MICRO_SERVER_ADDRESS=:50051 abjnet_prestation_service