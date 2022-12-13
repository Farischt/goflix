package config

import "fmt"

var (
	PORT        = GetIntEnv("PORT", ".env")
	DB_HOST     = GetEnv("POSTGRES_HOSTNAME_ORM", ".env.postgres")
	DB_PORT     = GetIntEnv("POSTGRES_PORT", ".env.postgres")
	DB_USER     = GetEnv("POSTGRES_USER", ".env.postgres")
	DB_PASSWORD = GetEnv("POSTGRES_PASSWORD", ".env.postgres")
	DB_NAME     = GetEnv("POSTGRES_DB", ".env.postgres")
	PSQL_URL    = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
)
