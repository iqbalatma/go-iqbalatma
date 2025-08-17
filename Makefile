ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

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


migrate:
	@echo "✅ Running all migration"
	migrate -path ./database/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" up

migrate\:rollback:
	@echo "✅ Rollback migration"
	migrate -path ./database/migration -database "$(DB_DRIVER)://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)" down



%:
	@: