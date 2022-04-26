gen:
	protoc -I=./proto/consignment --go_out=./proto/consignment/ --go_opt=paths=source_relative --go-grpc_out=./proto/consignment/ --go-grpc_opt=paths=source_relative consignment.proto

build:
	sudo docker build -t shippy-client .
run:
	sudo docker run shippy-client