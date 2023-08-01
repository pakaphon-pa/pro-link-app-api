package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
		Logger   LoggerConfig   `yaml:"logger"`
		Swagger  SwaggerConfig  `yaml:"swagger"`
		Redis    RedisConfig    `yaml:"redis"`
	}

	ServerConfig struct {
		Port string `yaml:"port,omitempty"`
	}

	DatabaseConfig struct {
		Host         string        `yaml:"host,omitempty"`
		Port         string        `yaml:"port,omitempty"`
		Username     string        `yaml:"username,omitempty"`
		Password     string        `yaml:"password,omitempty"`
		Name         string        `yaml:"name,omitempty"`
		SSLMode      string        `yaml:"sslmode,omitempty"`
		Timezone     string        `yaml:"timezone,omitempty"`
		MaxIdleConns int           `yaml:"maxIdleConns,omitempty"`
		MaxIdleTime  time.Duration `yaml:"maxIdleTime,omitempty"`
		MaxOpenConns int           `yaml:"maxOpenConns,omitempty"`
		MaxLifeTime  time.Duration `yaml:"maxLifeTime,omitempty"`
		SSLCert      string        `yaml:"sslcert,omitempty"`
		SSLKey       string        `yaml:"sslkey,omitempty"`
		SSLRootCert  string        `yaml:"sslrootcert,omitempty"`
	}

	RedisConfig struct {
		Dsn string `yaml:"dsn,omitempty"`
	}

	SwaggerConfig struct {
		Enable bool `yaml:"enable,omitempty"`
	}

	LoggerConfig struct {
		Level       string `mapstructure:"level"`
		OnCloud     bool   `mapstructure:"on-cloud"`
		Development bool   `mapstructure:"development"`
		Stacktrace  bool   `mapstructure:"stacktrace"`
		Caller      bool   `mapstructure:"caller"`
		DbLevel     string `mapstructure:"db-level"`
	}
)

const (
	defaultConfigName string = "config"
	defaultConfigType string = "yml"
	defaultConfigPath string = "../../etc"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", defaultConfigPath, "path to config file")
}

func LoadConfig(path string) *Config {
	if path != "" {
		configPath = path
	}

	var cfg = &Config{}
	workdir, err := os.Getwd()
	if err != nil {
		log.Panicln(err.Error())
	}

	filePath := filepath.Join(workdir, configPath)

	if err := os.Chdir(filePath); err != nil {
		panic(err)
	}

	viper.SetConfigName(defaultConfigName)
	viper.SetConfigType(defaultConfigType)
	viper.AddConfigPath(filePath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.MergeInConfig(); err != nil {
		log.Panicln(err.Error())
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicln(err.Error())
	}

	return cfg
}
