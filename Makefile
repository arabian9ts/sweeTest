PROJECT = sweeTest

.PHONY: go-cli
go-cli:
	@echo 'go get -v github.com/rubenv/sql-migrate/...'
	-@go get -v github.com/rubenv/sql-migrate/...
