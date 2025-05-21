package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AWSRegion    string
	S3Bucket     string
	S3Prefix     string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	AWSAccessKey string
	AWSSecretKey string
	Environment  string
}

func Load() (*Config, error) {
	// Setup viper to read env vars and optionally a .env file
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // automatically read environment variables

	// Optionally read from .env file if present in the current directory
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No .env file found or error reading it, continuing with env vars and defaults")
	}

	// Set default values
	viper.SetDefault("AWS_REGION", "us-east-1")
	viper.SetDefault("S3_BUCKET", "zuzu-reviews")
	viper.SetDefault("S3_PREFIX", "daily/")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "user")
	viper.SetDefault("DB_PASSWORD", "password")
	viper.SetDefault("DB_NAME", "reviews")
	viper.SetDefault("AWS_ACCESS_KEY_ID", "")
	viper.SetDefault("AWS_SECRET_ACCESS_KEY", "")
	viper.SetDefault("ENV", "development")

	cfg := &Config{
		AWSRegion:    viper.GetString("AWS_REGION"),
		S3Bucket:     viper.GetString("S3_BUCKET"),
		S3Prefix:     viper.GetString("S3_PREFIX"),
		DBHost:       viper.GetString("DB_HOST"),
		DBPort:       viper.GetString("DB_PORT"),
		DBUser:       viper.GetString("DB_USER"),
		DBPassword:   viper.GetString("DB_PASSWORD"),
		DBName:       viper.GetString("DB_NAME"),
		AWSAccessKey: viper.GetString("AWS_ACCESS_KEY_ID"),
		AWSSecretKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
		Environment:  viper.GetString("ENV"),
	}

	log.Println("Configuration loaded successfully.")
	return cfg, nil
}
