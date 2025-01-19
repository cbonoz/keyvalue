package constants

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	APP_NAME            string
	WEB_DOMAIN          string
	FROM_EMAIL          string
	SUPABASE_URL        string
	SUPABASE_ANON_KEY   string
	SUPABASE_JWT_SECRET string
	SERVER_PORT         string
	DATABASE_URL        string
	MIGRATION_FOLDER    string
	BREVO_API_KEY       string
	ENV                 string
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
	SUPABASE_JWT_SECRET = viper.GetString("SUPABASE_JWT_SECRET")
	DATABASE_URL = viper.GetString("GOOSE_DBSTRING")
	MIGRATION_FOLDER = viper.GetString("GOOSE_MIGRATION_DIR")
	WEB_DOMAIN = viper.GetString("WEB_DOMAIN")
	BREVO_API_KEY = viper.GetString("BREVO_API_KEY")
	FROM_EMAIL = viper.GetString("FROM_EMAIL")
	APP_NAME = viper.GetString("APP_NAME")
	SERVER_PORT = viper.GetString("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "8080"
	}

	if APP_NAME == "" {
		APP_NAME = "KeyValue API"
	}

	return nil
}
