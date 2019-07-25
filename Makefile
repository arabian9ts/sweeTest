PROJECT = sweeTest
VERSION = v0.0.1

.PHONY: all
all:
	make setup
	make test
	make build

.PHONY: docker-%
docker-%:
	make docker-build

.PHONY: docker-build
docker-build:
	@echo 'docker-compose up --build -d'
	-@docker-compose up --build -d

.PHONY: test
test:
	@echo 'Execute go test'
	-@go test github.com/arabian9ts/sweeTest/app/usecase/interactor github.com/arabian9ts/sweeTest/app/domain/model

.PHONY: setup
setup:
	@echo 'go get -u github.com/golang/dep/cmd/dep'
	-@go get -u github.com/golang/dep/cmd/dep
	@echo 'dep ensure'
	-@dep ensure

.PHONY: build
build:
	@echo 'go build -ldflags "-X main.version=$(VERSION)" -o sweeTest'
	-@go build -ldflags "-X main.version=$(VERSION)" -o sweeTest
