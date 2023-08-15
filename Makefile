BINARY_NAME:=sonarqube

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)


deps: ## Fetch any dependencies the project relies on
	go get go.uber.org/zap

## Build:
build: ## Build sonarqube and put the output binary in /bin/
	go build -o bin/${BINARY_NAME} main.go


run:
	go run main.go


compile: ## Compile for every platform
	echo "Compiling ${BINARY_NAME} for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/${BINARY_NAME}-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/${BINARY_NAME}-freebsd-386 main.go

clean: ## Remove binaries
	go clean
	rm -rf bin/

all: clean build


## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)