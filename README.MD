# Go Telebot Template

This project provides a template for initializing a Go bot using [telebot](https://github.com/tucnak/telebot), along with a structured directory for FSM, handlers, cache, middlewares, and database connections.

## Features

- Structured bot directory with FSM and handlers.
- Integrated with GORM, SQLite, Viper, and more.
- Automated setup with templates.
- Generates `.env` file for bot configuration.
- Modular setup for easy extension and maintenance.

## Requirements

- Go 1.23+

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/your-repo/go-telebot-init.git
    cd go-telebot-init
    ```

2. Run the `Makefile` to initialize the project:
    ```bash
    make all
    ```

## Commands in Makefile

### `make all`

Executes all the necessary steps:
- Set up environment variables.
- Initialize the project structure with templates.
- Install dependencies.
- Clean up template files.

### `make env`

Generates a `.env` file with the following environment variables:
- `TELEGRAM_TOKEN`: The bot token from Telegram.
- `SUPERUSER_ID`: The chat ID of the bot's superuser.
- `SUPERUSER_FIRSTNAME`: First name of the superuser.
- `SUPERUSER_LASTNAME`: Last name of the superuser.

### `make init`

Creates the necessary project structure, including directories and Go files for bot logic, FSM, handlers, middlewares, cache, database models, and utilities.

### `make deps`

Initializes the Go module, installs required Go packages, and tidies up the `go.mod` file.

### `make clean`

Removes the template files used during project initialization.

## After `make all`

Once you run the `make all` command, the project will be fully initialized and ready to use. It performs the following actions:

1. **Sets Up Environment Variables**: Creates a `.env` file with your bot's token and superuser details.
2. **Initializes Project Structure**: Generates the necessary directories and files for the bot's operation, FSM, handlers, middlewares, and database.
3. **Installs Dependencies**: Automatically installs required Go modules, ensuring your project has all dependencies.
4. **Cleans Up**: Removes template files after copying them into the appropriate directories.

### Running the Bot

After `make all` completes, you can start the bot by navigating to the `cmd/bot` directory and running:

```bash
go run main.go
```
The bot will be up and running with the initial configuration.

## Project Structure

```plaintext
cmd/
  └── bot/
      └── main.go
config/
  └── config.yaml
internal/
  └── configs/
      └── configs.go
pkg/
  ├── bot/
  │   ├── fsm/
  │   ├── handlers/
  │   │   ├── admin/
  │   │   └── public/
  │   └── middlewares/
  ├── cache/
  ├── database/
  │   ├── dbservice/
  │   └── models/
  ├── helpers/
  └── utils/
```

## Environment Setup

You'll be prompted to enter the following details:
- Bot Token (from [BotFather](https://core.telegram.org/bots#botfather))
- Superuser Chat ID
- Superuser First and Last Name

These details will be stored in a `.env` file automatically.

## Dependencies

The following Go libraries are used in this project:
- [GORM](https://gorm.io): For database interactions.
- [Viper](https://github.com/spf13/viper): For configuration management.
- [telebot.v3](https://gopkg.in/telebot.v3): For Telegram bot API.
- [eqmlog](https://github.com/lacolle87/eqmlog): Custom logging utility.

## License

This project is licensed under the MIT License.