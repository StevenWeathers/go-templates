package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ListenAddress string `mapstructure:"LISTEN_ADDRESS"`
}

func New(ctx context.Context, cfgFile string) *Config {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".webapp-example") // name of config file (without extension)
	viper.AddConfigPath("$HOME")           // adding home directory as first search path

	viper.SetDefault("LISTEN_ADDRESS", ":8080")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.WarnContext(ctx, fmt.Sprintf("Config file not found; ignore error if desired: %v", err))
		} else {
			logger.ErrorContext(ctx, fmt.Sprintf("Config file was found but another error was produced: %v", err))
		}
	}

	config := Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		logger.ErrorContext(ctx, fmt.Sprintf("unable to unmarshal config into struct, %v", err))
	}

	return &config
}
