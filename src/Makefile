createMongoDBcontainer:
	docker run -d -p 27017:27017 --name=MongoDBCRUD -v mongo_data:/data/db mongo_data

startContainer:
	docker start MongoDBCRUD

stopContainer:
	docker stop MongoDBCRUD

run:
	go run main.go