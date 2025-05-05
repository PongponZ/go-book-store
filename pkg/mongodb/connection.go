package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     *int
	Database string
}

func Connect(config *MongoConfig) (*mongo.Client, error) {
	var url string
	if config.Port == nil {
		url = fmt.Sprintf("mongodb+srv://%s:%s@%s/%s", config.Username, config.Password, config.Host, config.Database)
	} else {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin&authMechanism=SCRAM-SHA-1", config.Username, config.Password, config.Host, *config.Port, config.Database)
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
