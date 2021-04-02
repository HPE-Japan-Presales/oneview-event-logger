# Go Prams
GOCMD=go
GOBUILD=$(GOCMD) build -trimpath
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
CURRENT_VERSION=$(shell git describe --tags --abbrev=0)
BUILD_TARGET="cmd/main.go"
BUILD_PATH="bin/"
BUILD_BASE_NAME=oneview-event-logger

#Run Params
OV_ADDR="192.168.2.6"
OV_USER="golang"
OV_PASSWORD="golangtest"

test:
	@OV_ADDR=$(OV_ADDR) OV_USER=$(OV_USER) OV_PASSWORD=$(OV_PASSWORD) $(GORUN) $(BUILD_TARGET)

build:
	@$(GOCLEAN) all
	@echo Version:$(CURRENT_VERSION)
	@mkdir -p $(BUILD_PATH)
	@echo "== Build for Windows amd64"
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_PATH)$(BUILD_BASE_NAME) -ldflags "-X main.version=$(CURRENT_VERSION)" $(BUILD_TARGET)
	@tar zcvf $(BUILD_PATH)$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-windows-amd64.tar.gz -C $(BUILD_PATH) $(BUILD_BASE_NAME)
	@echo "== Build for OSX amd64"
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_PATH)$(BUILD_BASE_NAME) -ldflags "-X main.version=$(CURRENT_VERSION)" $(BUILD_TARGET)
	@tar zcvf $(BUILD_PATH)$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-darwin-amd64.tar.gz -C $(BUILD_PATH) $(BUILD_BASE_NAME)
	@echo "== Build for Linux amd64"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GOBUILD) -o $(BUILD_PATH)$(BUILD_BASE_NAME) -ldflags "-X main.version=$(CURRENT_VERSION)" $(BUILD_TARGET)
	@tar zcvf $(BUILD_PATH)$(BUILD_BASE_NAME)-$(CURRENT_VERSION)-linux-amd64.tar.gz -C $(BUILD_PATH) $(BUILD_BASE_NAME)
	@rm $(BUILD_PATH)$(BUILD_BASE_NAME)