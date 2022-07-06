# Authentication
protogen-microservice-authentication:
	@rm -Rf microservice-authentication/pkg/pbuff
	@protoc --proto_path=microservice-authentication/protobuf --go_out=plugins=grpc:. microservice-authentication/protobuf/*.proto
	@echo generating protbuff finished..

# Product
protogen-microservice-product:
	@rm -Rf microservice-product/pkg/pbuff
	@protoc --proto_path=microservice-product/protobuf --go_out=plugins=grpc:. microservice-product/protobuf/*.proto
	@echo generating protbuff finished..

# Order
protogen-microservice-order:
	@rm -Rf microservice-order/pkg/pbuff
	@protoc --proto_path=microservice-order/protobuf --go_out=plugins=grpc:. microservice-order/protobuf/*.proto
	@echo generating protbuff finished..

# Deploy
dockerize-microservice-authentication:
	@docker build --build-arg server=microservice-authentication -f ./microservice-authentication/Dockerfile -t microservice-authentication:latest .
	@docker images --filter=reference='microservice-authentication:latest'

dockerize-microservice-product:
	@docker build --build-arg server=microservice-product -f ./microservice-product/Dockerfile -t microservice-product:latest .
	@docker images --filter=reference='microservice-product:latest'

dockerize-microservice-order:
	@docker build --build-arg server=microservice-order -f ./microservice-order/Dockerfile -t microservice-order:latest .
	@docker images --filter=reference='microservice-order:latest'

dockerize-all-service: dockerize-microservice-authentication dockerize-microservice-product dockerize-microservice-order

start-gocommerce-docker:
	@docker-compose up -d

stop-gocommerce-docker:
	@docker-compose down

test-run:
	go test gocommerce/microservice-authentication/services
	go test gocommerce/microservice-order/services
	go test gocommerce/microservice-product/services

test-build-run: test-run dockerize-all-service start-gocommerce-docker

build-stop: stop-gocommerce-docker