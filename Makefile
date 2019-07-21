PROJECT = sweeTest

.PHONY: go-cli
go-cli:
	@echo 'go get -v github.com/rubenv/sql-migrate/...'
	-@go get -v github.com/rubenv/sql-migrate/...

.PHONY: test
test:
	@echo 'Execute go test'
	-@go test github.com/arabian9ts/sweeTest/app/interactor github.com/arabian9ts/sweeTest/app/domain/model
