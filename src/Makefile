# NOTE: For some reason, BUILD_DATE assignment doesn't work in Windows + Visual Studio Code: make hangs at this point

BUILD_DATE := $(shell date "+%Y-%m-%dT%H:%M:%S")

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
	@go run main.go $(BUILD_DATE) - 112
