




mysqlup:
	@docker run --name local-mysql -e MYSQL_ROOT_PASSWORD=sqlpass -d -p 3306:3306  mysql:latest



watch:
	go run examples/main.go