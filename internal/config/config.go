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
		Server    ServerConfig   `mapstructure:"server"`
		Database  DatabaseConfig `mapstructure:"database"`
		Logger    LoggerConfig   `mapstructure:"logger"`
		Swagger   SwaggerConfig  `mapstructure:"swagger"`
		Redis     RedisConfig    `mapstructure:"redis"`
		JwtConfig JwtConfig      `mapstructure:"jwt"`
		Email     EmailConfig    `mapstructure:"email"`
	}

	ServerConfig struct {
		Port         string `mapstructure:"port,omitempty"`
		ClientOrigin string `mapstructure:"client-origin,omitempty"`
	}

	DatabaseConfig struct {
		Host         string        `mapstructure:"host,omitempty"`
		Port         string        `mapstructure:"port,omitempty"`
		Username     string        `mapstructure:"username,omitempty"`
		Password     string        `mapstructure:"password,omitempty"`
		Name         string        `mapstructure:"name,omitempty"`
		SSLMode      string        `mapstructure:"sslmode,omitempty"`
		Timezone     string        `mapstructure:"timezone,omitempty"`
		MaxIdleConns int           `mapstructure:"maxIdleConns,omitempty"`
		MaxIdleTime  time.Duration `mapstructure:"maxIdleTime,omitempty"`
		MaxOpenConns int           `mapstructure:"maxOpenConns,omitempty"`
		MaxLifeTime  time.Duration `mapstructure:"maxLifeTime,omitempty"`
		SSLCert      string        `mapstructure:"sslcert,omitempty"`
		SSLKey       string        `mapstructure:"sslkey,omitempty"`
		SSLRootCert  string        `mapstructure:"sslrootcert,omitempty"`
	}

	RedisConfig struct {
		Dsn string `mapstructure:"dsn,omitempty"`
	}

	SwaggerConfig struct {
		Enable bool `mapstructure:"enable,omitempty"`
	}

	LoggerConfig struct {
		Level       string `mapstructure:"level"`
		OnCloud     bool   `mapstructure:"on-cloud"`
		Development bool   `mapstructure:"development"`
		Stacktrace  bool   `mapstructure:"stacktrace"`
		Caller      bool   `mapstructure:"caller"`
		DbLevel     string `mapstructure:"db-level"`
	}

	JwtConfig struct {
		AccessSecret   string `mapstructure:"access-secret"`
		AccessExpired  string `mapstructure:"access-expired"`
		AccessMaxAge   int    `mapstructure:"access-max-age"`
		RefreshSecret  string `mapstructure:"refresh-secret"`
		RefreshExpired string `mapstructure:"refresh-expired"`
		RefreshMaxAge  int    `mapstructure:"refresh-max-age"`
	}

	EmailConfig struct {
		Host     string              `mapstructure:"host"`
		Port     int                 `mapstructure:"port"`
		Username string              `mapstructure:"username"`
		Password string              `mapstructure:"password"`
		From     string              `mapstructure:"from"`
		Template TemplateEmailConfig `mapstructure:"template"`
	}

	TemplateEmailConfig struct {
		ConfirmAccount string `mapstructure:"confirm-account"`
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
