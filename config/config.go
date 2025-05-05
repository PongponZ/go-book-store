package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDB MongoDBConfig
	HTTP    HTTPConfig
}

type MongoDBConfig struct {
	Host           string
	Port           *int
	Username       string
	Password       string
	Database       string
	UserCollection string
}

type HTTPConfig struct {
	Port int
}

const (
	MONGO_HOST            = "MONGO_HOST"
	MONGO_PORT            = "MONGO_PORT"
	MONGO_USERNAME        = "MONGO_USERNAME"
	MONGO_PASSWORD        = "MONGO_PASSWORD"
	MONGO_DATABASE        = "MONGO_DATABASE"
	MONGO_USER_COLLECTION = "MONGO_USER_COLLECTION"

	HTTP_PORT = "HTTP_PORT"
)

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	mongoConfig := getMongoDBConfig()
	httpConfig := getHTTPConfig()

	return &Config{
		MongoDB: *mongoConfig,
		HTTP:    *httpConfig,
	}
}

func getHTTPConfig() *HTTPConfig {
	httpPortInt, err := strconv.Atoi(os.Getenv(HTTP_PORT))
	if err != nil {
		log.Fatal("error converting HTTP_PORT to int")
	}

	return &HTTPConfig{
		Port: httpPortInt,
	}
}

func getMongoDBConfig() *MongoDBConfig {
	mongoHost := os.Getenv(MONGO_HOST)
	mongoPortStr := os.Getenv(MONGO_PORT)
	mongoUsername := os.Getenv(MONGO_USERNAME)
	mongoPassword := os.Getenv(MONGO_PASSWORD)
	mongoDatabase := os.Getenv(MONGO_DATABASE)
	mongoUserCollection := os.Getenv(MONGO_USER_COLLECTION)

	var mongoPortInt *int
	if mongoPortStr != "" {
		port, err := strconv.Atoi(mongoPortStr)
		if err != nil {
			log.Fatal("error converting MONGO_PORT to int")
		}
		mongoPortInt = &port
	}

	return &MongoDBConfig{
		Host:           mongoHost,
		Port:           mongoPortInt,
		Username:       mongoUsername,
		Password:       mongoPassword,
		Database:       mongoDatabase,
		UserCollection: mongoUserCollection,
	}
}
