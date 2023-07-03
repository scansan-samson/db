




mysqlup:
	@docker run --name local-mysql -e MYSQL_ROOT_PASSWORD=sqlpass -d -p 3306:3306  mysql:latest

fmt:
	@go mod tidy
	@goimports -w .
	@gofmt -w -s .
	@go clean ./...

watch:
	go run examples/main.go


test:
	go test -v -coverprofile=profile.cov ./...

