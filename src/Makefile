# BUILD_DATE := $(shell date "+%Y-%m-%dT%H:%M:%S")
BUILD_DATE := $("2024-12-29")
BUILD_NUMBER := $("112")

createMongoDBcontainer:
	docker run -d -p 27017:27017 --name=Open311ToGo -v mongo_data:/data/db mongo_data

startContainer:
	docker start Open311ToGo

stopContainer:
	docker stop Open311ToGo

build:
	echo "Building the application with BuildDate=$(BUILD_DATE)"
	@go build -v -ldflags "-X main.BuildDate=$(BUILD_DATE)" -o open311togo ../build

run:
	@go run main.go $(BUILD_DATE) $(BUILD_NUMBER)
