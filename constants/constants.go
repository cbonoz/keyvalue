package constants

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	SUPABASE_URL      string
	SUPABASE_ANON_KEY string
	DATABASE_URL string
	BREVO_API_KEY string
	ENV               string
)

func Init() error {
	// Set default environment to development
	ENV = "development"
	if envVar := viper.GetString("ENV"); envVar != "" {
		ENV = envVar
	}

	// Setup viper
	viper.SetConfigType("env")

	// Try loading .env.prod in production, fallback to .env
	if ENV == "production" {
		viper.SetConfigName(".env.prod")
	} else {
		viper.SetConfigName(".env")
	}

	viper.AddConfigPath(".")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	// Set constants from viper
	SUPABASE_URL = viper.GetString("SUPABASE_URL")
	SUPABASE_ANON_KEY = viper.GetString("SUPABASE_ANON_KEY")
	DATABASE_URL = viper.GetString("DATABASE_URL")
	BREVO_API_KEY = viper.GetString("BREVO_API_KEY")

	return nil
}

