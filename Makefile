SHELL := /bin/bash
.PHONY: env init clean deps all

MODULE=$(shell basename $(shell pwd))

all: env init deps clean

env:
	@read -p "Enter your bot token: " token; \
	read -p "Enter your SuperUser Chat ID: " chat_id; \
	read -p "Enter SuperUser First Name: " first_name; \
	read -p "Enter SuperUser Last Name: " last_name; \
	echo "BOT_TOKEN=$$token" > .env; \
	echo "SU_CHAT_ID=$$chat_id" >> .env; \
	echo "SUPERUSER_NAME=SU" >> .env; \
	echo "SUPERUSER_FIRSTNAME=$$first_name" >> .env; \
	echo "SUPERUSER_LASTNAME=$$last_name" >> .env; \
	echo ".env file created successfully."

init:
	@mkdir -p cmd/bot/ internal/ configs/ pkg/cache pkg/helpers pkg/utils
	@mkdir -p pkg/bot/fsm pkg/bot/handlers/admin pkg/bot/handlers/public pkg/bot/middlewares
	@mkdir -p pkg/database/models pkg/database/dbservice

	@cp templates/cmd/bot/main_template.go cmd/bot/main.go
	@cp templates/config/config_template.yaml config/config.yaml
	@cp templates/internal/configs/configs_template.go internal/configs/configs.go

	@cp templates/pkg/bot/bot_template.go pkg/bot/bot.go
	@cp templates/pkg/bot/fsm/fsm_template.go pkg/bot/fsm/fsm.go
	@cp templates/pkg/bot/handlers/admin/admin_template.go pkg/bot/handlers/admin/admin.go
	@cp templates/pkg/bot/handlers/public/public_template.go pkg/bot/handlers/public/public.go
	@cp templates/pkg/bot/middlewares/middlewares_template.go pkg/bot/middlewares/middlewares.go

	@cp templates/pkg/database/database_template.go pkg/database/database.go
	@cp templates/pkg/database/models/models_template.go pkg/database/models/models.go
	@cp templates/pkg/database/dbservice/dbservice_template.go pkg/database/dbservice/dbservice.go

	@cp templates/pkg/cache/cache_template.go pkg/cache/cache.go

	@cp templates/pkg/helpers/helpers_template.go pkg/helpers/helpers.go
	@cp templates/pkg/utils/utils_template.go pkg/utils/utils.go
	@find . -type f -name '*.go' -exec sed -i.bak "s|go-telebot-init|$(MODULE)|g" {} +
	@echo "Initialization complete."

deps:
	go mod init $(MODULE)
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite
	go get -u github.com/spf13/viper
	go get -u gopkg.in/telebot.v3
	go get -u github.com/lacolle87/eqmlog
	go mod tidy

clean:
	@rm -rf templates/
	@echo "Template files cleaned."


