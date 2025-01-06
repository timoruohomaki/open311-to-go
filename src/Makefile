# NOTE: For some reason, BUILD_DATE assignment doesn't work in Windows + VSCode: make extension hangs at this point

# BUILD_DATE := $(shell date "+%Y-%m-%dT%H:%M:%S")
BUILD_DATE := $("1")
BUILD_NUMBER := $("0")
BUILD_ENV := $("DEV")

createMongoDBcontainer:
	docker run -d -p 27017:27017 --name=Open311ToGo -v mongo_data:/data/db mongo_data

startContainer:
	docker start Open311ToGo

stopContainer:
	docker stop Open311ToGo

build:
	
	# @go build -v -ldflags "-X main.BuildDate=$(BUILD_DATE)" -o open311togo ../build

run:
	@BUILD_DATE=$(BUILD_DATE) BUILD_NUMBER="113" BUILD_ENV="DEV" go run main.go

winrun:
	@go run main.go