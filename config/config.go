package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddr  string `mapstructure:"SERVER_ADDR" validate:"required"`
	DBHost      string `mapstructure:"DB_HOST" validate:"required"`
	DBName      string `mapstructure:"DB_NAME" validate:"required"`
	DBUser      string `mapstructure:"DB_USER" validate:"required"`
	DBPort      string `mapstructure:"DB_PORT" validate:"required"`
	DBPaassword string `mapstructure:"DB_PASSWORD" validate:"required"`

	// TwilioAuthToken  string `mapstructure:"TWILIO_AUTHTOKEN" validate:"required"`
	// TwilioAccountSID string `mapstructure:"TWILIO_ACCOUNTSID" validate:"required"`
	// TwilioServiceSID string `mapstructure:"TWILIO_SERVICESID" validate:"required"`
	Key string `mapstructure:"KEY" validate:"required"`
	// RazorpayKey    string `mapstructure:"razorpay_key" validate:"required"`
	// RazorpaySecret string `mapstructure:"razorpay_secret"`
	// // KeyAdmin         string `mapstructure:"KEY_ADMIN" validate:"required"`
	// PayKeyID         string `mapstructure:"KEY_ID_FOR_PAY" validate:"required"`
	// PaySecretKey     string `mapstructure:"SECRET_KEY_FOR_PAY" validate:"required"`
}

func LoadConfig() (Config, error) {
	var config Config
	viper.SetConfigFile("../.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading the config file: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
