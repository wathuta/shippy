gen:
	protoc --proto_path=. --go_out=./proto/vessel --micro_out=./proto/vessel proto/vessel/vessel.proto 
build:
	sudo docker build -t vessel-service .
run:
	sudo docker run vessel-service