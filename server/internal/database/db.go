package database

import (
	"context"
	"fmt"
	"server/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDatabase(cfg config.Config) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%d", cfg.Database.Host, cfg.Database.Port)
	clientOpts := options.Client().ApplyURI(uri)

	var client *mongo.Client
	var err error

	for i := 0; i < cfg.Database.Retries; i++ {
		client, err = mongo.Connect(context.TODO(), clientOpts)
		if err == nil {
			break
		}
	}
	return client.Database(cfg.Database.Db), err
}
