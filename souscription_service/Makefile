compile_proto:
	protoc  --go_out=. --micro_out=. proto/souscription/souscription.proto
build: 
		docker build -t abjnet_souscription_service .
run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 abjnet_souscription_service