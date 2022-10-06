micro:
	protoc --proto_path=./ --go_out=./proto/consignment --micro_out=./proto/consignment  proto/consignment/consignment.proto
vessel:
	protoc --proto_path=./ --go_out=./proto/vessel --micro_out=./proto/vessel proto/vessel/vessel.proto
build:
	sudo docker build --build-arg buildtime_variable="mongodb+srv://brayo:brayo@myblogcluster.976g4.mongodb.net/shippy?retryWrites=true&w=majority"  -t consignment-service .
run:
	sudo docker run -p 50051:50051 -e MICRO_REGISTRY=mdns -e MICRO_SERVER_ADDRESS=:50051 -e DB_HOST='mongodb+srv://brayo:brayo@myblogcluster.976g4.mongodb.net/shippy?retryWrites=true&w=majority' -e DISABLE_AUTH=true --name=consignment-service consignment-service
	