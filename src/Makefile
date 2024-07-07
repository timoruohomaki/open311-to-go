createMongoDBcontainer:
	docker run -d -p 27017:27017 --name=MongoDBCRUD -v mongo_data:/data/db mongo_data

startContainer:
	docker start MongoDBCRUD

stopContainer:
	docker stop MongoDBCRUD

build:
	go build -v -o ../build

run:
	go run main.go
