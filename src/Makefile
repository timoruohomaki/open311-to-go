# BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

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
	echo "Running the application with BuildDate set as $(BUILD_DATE)"
	@go run main.go
