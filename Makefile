ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

dev:
	air -c .air.toml

serve:
	go run main.go

seed:
	go run main.go seeder


make\:migration:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "⚠️  Use this format : make make:migration create_users_table"; \
	else \
		migrate create -ext sql -dir "database/migration" $(filter-out $@,$(MAKECMDGOALS)); \
		echo "📦 New migration created: $(filter-out $@,$(MAKECMDGOALS))"; \
	fi


make\:seeder:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "⚠️  Use this format : make make:seeder user_seeder"; \
	else \
	  	mkdir database/seeder; \
	  	touch ./database/seeder/$(filter-out $@,$(MAKECMDGOALS)).go;\
		echo "📦 New seeder created: $(filter-out $@,$(MAKECMDGOALS))"; \
	fi

make\:model:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "⚠️  Use this format : make make:model user"; \
	else \
	  	mkdir app/model; \
	  	touch ./app/model/$(filter-out $@,$(MAKECMDGOALS)).go;\
		echo "📦 New model created: $(filter-out $@,$(MAKECMDGOALS))"; \
	fi


make\:controller:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "⚠️  Use this format : make make:controller management/user_controller"; \
	else \
		FILE=$(filter-out $@,$(MAKECMDGOALS)); \
		DIR=app/controller/$$(dirname $$FILE); \
		mkdir -p $$DIR; \
		touch app/controller/$$FILE.go; \
		echo "📦 New controller created: app/controller/$$FILE.go"; \
	fi

make\:service:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "⚠️  Use this format : make make:service management/user_service"; \
	else \
		FILE=$(filter-out $@,$(MAKECMDGOALS)); \
		DIR=app/service/$$(dirname $$FILE); \
		mkdir -p $$DIR; \
		touch app/service/$$FILE.go; \
		echo "📦 New service created: app/service/$$FILE.go"; \
	fi

make\:repository:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "⚠️  Use this format : make make:repository management/user_service"; \
	else \
		FILE=$(filter-out $@,$(MAKECMDGOALS)); \
		DIR=app/repository/$$(dirname $$FILE); \
		mkdir -p $$DIR; \
		touch app/repository/$$FILE.go; \
		echo "📦 New repository created: app/repository/$$FILE.go"; \
	fi

migrate:
	@echo "✅ Running all migration"
	migrate -path ./database/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up

migrate\:rollback:
	@echo "✅ Rollback migration"
	migrate -path ./database/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down



%:
	@: