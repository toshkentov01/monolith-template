#user service
protoc \
		--proto_path=/usr/local/include \
		--proto_path=./protos/user-service \
		--gofast_out=plugins=grpc:. \
		--gorm_out=. \
		service.proto
