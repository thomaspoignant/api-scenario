# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet

GOFMT=gofmt
GOLINT=golint
BINARY_NAME=rest-scenario
CMD_FOLDER=cmd
COVERAGE_FOLDER=.coverage

all: generate build lint test

clean:
	$(GOCMD) clean -i .

clean-all: clean ## remove all generated artifacts and clean all build artifacts
	rm -rf docs vendor $(COVERAGE_FOLDER) $(BINARY_NAME)
	find . -name "*_enumer.go" -delete # remove generated files

update-dependencies: ## update golang dependencies
	go mod vendor

generate:
	$(GOGET) github.com/alvaroloes/enumer
	$(GOCMD) generate ./...

test: update-dependencies generate
	$(GOTEST) -short -mod=vendor ./...

build: update-dependencies generate
	$(GOBUILD) -mod=vendor -o rest-scenario $(CMD_FOLDER)/rest-scenario/main.go

coverage:
	mkdir -p .coverage/
	$(GOTEST) -short -mod=vendor -coverprofile=.coverage/profile.cov ./...

lint:
	$(GOGET) golang.org/x/lint/golint
	$(GOLINT) -set_exit_status $($(GOCMD) list ./... | grep -v /vendor/)
