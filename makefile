gen:
	rm -rf idl/pb
	mkdir -p idl/pb
	protoc -I=idl/proto -I=../googleapis/\
        --go_out=idl/pb --go_opt=paths=source_relative \
        --go-grpc_out=idl/pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=idl/pb --grpc-gateway_opt=paths=source_relative \
        idl/proto/link/*
