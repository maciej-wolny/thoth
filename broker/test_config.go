package main

import (
	"os"
)

type DatabaseConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

var defaultDatabaseConfig = &DatabaseConfig{
	host:     "localhost",
	port:     "5432",
	user:     "helsing_edge_user",
	password: "123",
	dbname:   "helsing_edge",
}

func GetDatabaseTestConfig() *DatabaseConfig {
	return &DatabaseConfig{
		host:     getEnv("DB_HOST", defaultDatabaseConfig.host),
		port:     getEnv("DB_PORT", defaultDatabaseConfig.port),
		user:     getEnv("DB_USER", defaultDatabaseConfig.user),
		password: getEnv("DB_PASS", defaultDatabaseConfig.password),
		dbname:   getEnv("DB_NAME", defaultDatabaseConfig.dbname),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
