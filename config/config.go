package config

import (
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var (
	cfg  *Config
	once = &sync.Once{}
)

type Config struct {
	App App
	DB  DB
}

type App struct {
	Port int
}

type DB struct {
	Dsn string
}

func InitConfig(configDir string) *Config {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(configDir)

		if err := viper.ReadInConfig(); err != nil {
			logger.Fatal().Err(err).Msg("failed to read config")
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			logger.Fatal().Err(err).Msg("failed to unmarshal config")
		}

		viper.OnConfigChange(func(e fsnotify.Event) {

			if err := viper.Unmarshal(&cfg); err != nil {
				logger.Error().Err(err).Msg("failed to unmarshal config")
			}

			logger.Info().Msg("config file changed and updated")
		})

		viper.WatchConfig()
	})

	return cfg
}

// Get returns an up-to-date copy of the config
func Get() Config {
	return *cfg
}
