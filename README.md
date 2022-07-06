## gocommerce Ery Ansari Siregar
build microservices with solutions stack with :
1. Golang
2. MongoDB
3. GRPC
4. Docker
5. Mono Repo With 3 Microservice
##

###Script Docker in Makefile
```bash
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
```
##

### Simplify Test, Build & Run Docker
```bash
make test-build-run
```

```
erysiregar@ery gocommerce % make test-build-run
go test gocommerce/microservice-authentication/services
ok      gocommerce/microservice-authentication/services (cached)
go test gocommerce/microservice-order/services
ok      gocommerce/microservice-order/services  0.829s
go test gocommerce/microservice-product/services
ok      gocommerce/microservice-product/services        (cached)
[+] Building 24.7s (18/18) FINISHED
```

```
[+] Running 5/5
 ⠿ Container dbtwo                        Running                                                                                                                                                                                  0.0s
 ⠿ Container dbone                        Running                                                                                                                                                                                  0.0s
 ⠿ Container microservice-order           Started                                                                                                                                                                                  2.1s
 ⠿ Container microservice-product         Running                                                                                                                                                                                  0.0s
 ⠿ Container microservice-authentication  Running  
```
##

### Stop Run Docker
```bash
make build-stop
```
##

### Usage Insomnia
```
for download : https://insomnia.rest/download
for using : https://docs.insomnia.rest/insomnia/grpc
```
##

### Import Collection to Insomnia from root project
name file of collection insomnia:
```
COLLECTION-Insomnia_2022-07-06.json
```
##

### System Diagram
![diagram](https://i.postimg.cc/cCbZs8nz/diagram.png)

##### Microservice-authentication
#####Login ![login](https://i.postimg.cc/VvnQcHzJ/auth-login.png)
```bash
after login, use accessToken, and put in to Header with key Authorization : Bearer {accessToken}
{"Authorization": "Bearer accessToken"}
```
###
#####Register ![register](https://i.postimg.cc/Qx75Rg5v/auth-register.png)
##

##### Microservice-product
#####Create Product ![create](https://i.postimg.cc/6p0cyscd/product-create.png)
####
#####Get List Product ![get-list-product](https://i.postimg.cc/sDP9NFtt/product-list.png)
####
#####Get Product By ID ![get-product-byID](https://i.postimg.cc/VkBGn0f0/product-by-ID.png)
##

##### Microservice-order
#####Create Order ![create](https://i.postimg.cc/BnYFrJNw/order-create.png)
####
#####Get Order By User ID ![get-order-byUserID](https://i.postimg.cc/HLrMHFhW/order-by-User-ID.png)
####
#####Get Order By ID ![get-order-byID](https://i.postimg.cc/0jPr64cw/order-by-ID.png)
##


```
Notes : if asset not load perfectly, you can get it in folder V-README-asset
```