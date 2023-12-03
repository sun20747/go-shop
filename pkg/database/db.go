package database

import (
	"github.com/go-shop/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func DbConn(cfg *config.Config) *mongo.Client {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	return nil
}
