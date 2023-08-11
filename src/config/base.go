package config

import (
	"os"
)

var (
	ENVIRONMENT      string
	SECRET_KEY_TOKEN string
	API_KEYS         interface{}
	API_VERSION      = "v1"
	UNIACCOUNT_NAME  = "uniaccount"
	TOKEN_TYPE       = "JWT"

	PORT         = 3001
	DB_INIT_PORT = 3100
	DEBUG        = true
	HOST         = "0.0.0.0"

	//PostgreSQL
	URL_POSTGRESQL = ""
)

func LoadEnvs() {

	// Obtener valores de las variables de entorno
	ENVIRONMENT = getEnv("ENVIRONMENT")
	SECRET_KEY_TOKEN = getEnv("SECRET_KEY_TOKEN")

	URL_POSTGRESQL = getEnv("URL_POSTGRESQL")

}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		switch key {
		case "ENVIRONMENT":
			return "local"
		case "SECRET_KEY_TOKEN":
			return "default_secret_key"
		case "UNIACCOUNTS_API_KEY":
			return "api_key"
		case "UNIACCOUNTS_API_SERVICE_NAME":
			return "service"
		case "SMTP_USER":
			return "user@mail.com"
		case "SMTP_PASSWORD":
			return "password"
		default:
			return ""
		}
	}
	return value
}
