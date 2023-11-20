http:
	@echo "Running Go Http Application"
	go run main.go http

migrate:
	@echo "Migrate Table to Database"
	cd db/migrate && go run migrate.go