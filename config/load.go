package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"log"
	"os"
)


// Config ...
type Config struct {
	PostgresUser string
	PostgresPassword string
	PostgresDBName string
	Unidentified int64
	Identified int64
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	c := Config{}

	c.PostgresUser=cast.ToString(getOrReturnDefault("POSTGRES_USER","user"))
	c.PostgresPassword=cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD","password"))
	c.PostgresDBName=cast.ToString(getOrReturnDefault("POSTGRES_DB","wallet"))
	c.Unidentified = cast.ToInt64(getOrReturnDefault("UNIDENTIFIED",10000))
	c.Identified = cast.ToInt64(getOrReturnDefault("IDENTIFIED",100000))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}