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
	@mkdir -p config/ cmd/internal pkg/bot/router pkg/bot/middlewares  pkg/utils
	@mkdir -p pkg/database/models pkg/database/dbservice
	@cp templates/configs/config_template.yaml config/config.yaml
	@cp templates/main_template.go cmd/bot/main.go
	@cp templates/config_template.go internal/config.go
	@cp templates/bot_template.go pkg/bot/bot.go
	@cp templates/router_template.go pkg/bot/router/router.go
	@cp templates/middlewares_template.go pkg/bot/middlewares/middlewares.go
	@cp templates/database_template.go pkg/database/database.go
	@cp templates/models_template.go pkg/database/models/models.go
	@cp templates/dbservice_template.go pkg/database/dbservice/dbservice.go
	@cp templates/utils_template.go pkg/utils/utils.go
	@find . -type f -name '*.go' -exec sed -i '' "s|go-telebot-init|$(MODULE)|g" {} +
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


