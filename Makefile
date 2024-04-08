build:
	@echo "\033[0;32mBuilding the application...\033[0m"
	@touch .env
	@echo DB_PASSWORD="''" > .env
	@docker-compose build todo-app
	@echo "\033[0;32m  NOTIFY : Select your DB_PASSWORD in .env file... \033[0m"
	
	

run:
	@echo "\033[0;33mRunning the application...\033[0m"
	@docker-compose up todo-app

migrate:
	@echo "\033[0;35mRunning migrations...\033[0m"
	@migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	@echo "\033[0;36mInitializing Swagger...\033[0m"
	@swag init -g cmd/main.go