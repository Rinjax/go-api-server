package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Api     ApiServerConfig `json:"api"`
	App AppConfig `json:"app"`
	Auth AuthConfig `json:"auth"`
	Hashing HashingConfig  `json:"hashing"`
	Store StoreConfig `json:"store"`
	Log LogConfig `json:"log"`
}

type AppConfig struct {
	Name string
}

type AuthConfig struct {
	SigningKey string `mapstructure:"signing_key"`
}

type ApiServerConfig struct {
	Addr string `json:"addr"`
}

type HashingConfig struct {
	Hasher string         `json:"hasher"`
	Argon2 Argon2iConfig `json:"argon2"`
}

type Argon2iConfig struct {
	TimeCost    uint32 `mapstructure:"time_cost"`
	MemoryCost  uint32 `mapstructure:"memory_cost"`
	Parallelism uint8  `json:"parallelism"`
	HashLength  uint32 `mapstructure:"hash_length"`
}

type StoreConfig struct {
	Storer string `json:"storer"`
}

type LogConfig struct {
	Logger string `json:"logger"`
	Level string `json:"level"`
}

func GetConfig() (*Config, error) {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func getConfigPath(env string) string {
	// todo switch on env 
	return "config/config"
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	cfg := &Config{}

	err := v.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", cfg)

	return cfg, nil
}
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	//v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}

		return nil, err
	}

	return v, nil
}