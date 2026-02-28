package config

import "os"

type Config struct {
	Port  string
	DbDsn string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Example:
	// host=localhost user=postgres password=postgres dbname=agnos port=5432 sslmode=disable TimeZone=Asia/Bangkok
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=agnos port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	}

	return Config{
		Port:  port,
		DbDsn: dsn,
	}
}
