help:
	@echo "$$HELP_TEXT"

dependencies:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz migrate && sudo mv migrate /usr/bin

build:
	sudo docker compose build

run:
	sudo docker compose up 
	# --abort-on-container-exit

run_build:
	sudo docker compose down
	sudo docker compose up --build

down:
	sudo docker compose down

action=up
n=1
migrate:
	migrate -path ./migration/todoService -database "postgres://todo_admin:123456789@0.0.0.0:5432/todo?sslmode=disable" -verbose $(action) $(n)
	migrate -path ./migration/authService -database "postgres://auth_admin:123456789@0.0.0.0:5432/auth?sslmode=disable" -verbose $(action) $(n)

migrate_create:
	migrate create -ext sql -dir ./migration/todoService -seq category_schema
	migrate create -ext sql -dir ./migration/todoService -seq todo_schema
	migrate create -ext sql -dir ./migration/authService -seq accounts_schema

define HELP_TEXT
Use the following commands:
	make dependencies
		install dependencies needed for execution of this project

	make build
		build the project binary

	make migrate [up|down|force] N
		run migrations. default is ACTION=up with N=1

	migrate_create
		create migrations files
		
	make run
		run project using docker compose
	
	make run_build
		build project and run using docker compose
	
	make down
		down all started containers in docker compose

	make help
		print this manual

endef

export HELP_TEXT
