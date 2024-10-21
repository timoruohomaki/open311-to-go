createMongoDBcontainer:
	docker run -d -p 27017:27017 --name=Open311ToGo -v mongo_data:/data/db mongo_data

startContainer:
	docker start Open311ToGo

stopContainer:
	docker stop Open311ToGo

build:
	@go build -v -o ../build

run:
	@go run main.go
