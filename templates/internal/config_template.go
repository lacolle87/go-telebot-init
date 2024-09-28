package configs

import (
	"fmt"
	"github.com/lacolle87/eqmlog"
	"github.com/spf13/viper"
	"log/slog"
)

func Load() error {
	if err := loadEnv(); err != nil {
		return err
	}
	if err := loadConfig(); err != nil {
		return err
	}
	if err := setupLogger(); err != nil {
		return err
	}
	return nil
}

func loadEnv() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func loadConfig() error {
	viper.SetConfigFile("configs/config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.MergeInConfig(); err != nil {
		return fmt.Errorf("error reading YAML loader file: %w", err)
	}
	return nil
}

func setupLogger() error {
	multiWriter := eqmlog.LoadLogger()
	logger := slog.New(slog.NewTextHandler(multiWriter, nil))
	slog.SetDefault(logger)
	return nil
}
