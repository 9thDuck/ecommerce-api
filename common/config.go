package common

import (
	"time"
)

type AppConfig struct {
	postgresDbDSN                      string
	jwtSecret                          string
	expiryAccessTokenDurationInMinutes time.Duration
	expiryRefreshTokenDurationInHours  time.Duration
}

func (config *AppConfig) GetPostgresDbDsn() string {
	return config.postgresDbDSN
}

func (config *AppConfig) GetJwtSecret() string {
	return config.jwtSecret
}

func (config *AppConfig) GetExpiryAccessTokenDurationInMinutes() time.Duration {
	return config.expiryAccessTokenDurationInMinutes
}

func (config *AppConfig) GetExpiryRefreshTokenDurationInHours() time.Duration {
	return config.expiryRefreshTokenDurationInHours
}

var APP_CONFIG AppConfig = AppConfig{}

func (config *AppConfig) ValidateAndSetup() {

	POSTGRES_DB_DSN := GetEnvOrThrow("POSTGRES_DB_DSN")
	JWT_SECRET := GetEnvOrThrow("JWT_SECRET")

	EXPIRY_ACCESS_TOKEN_DURATION_IN_MINUTES, err := EnvVarToTimeDuration("EXPIRY_ACCESS_TOKEN_DURATION_IN_MINUTES", time.Minute)
	if err != nil {
		LogFatalCustomError("env variable EXPIRY_ACCESS_TOKEN_DURATION_IN_MINUTES set with invalid value", err)
	}

	EXPIRY_REFRESH_TOKEN_DURATION_IN_HOURS, err := EnvVarToTimeDuration("EXPIRY_REFRESH_TOKEN_DURATION_IN_HOURS", time.Hour)
	if err != nil {
		LogFatalCustomError("env variable EXPIRY_REFRESH_TOKEN_DURATION_IN_HOURS set with invalid value", err)
	}

	config.postgresDbDSN = POSTGRES_DB_DSN
	config.jwtSecret = JWT_SECRET
	config.expiryAccessTokenDurationInMinutes = EXPIRY_ACCESS_TOKEN_DURATION_IN_MINUTES
	config.expiryRefreshTokenDurationInHours = EXPIRY_REFRESH_TOKEN_DURATION_IN_HOURS
}
