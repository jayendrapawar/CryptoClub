package store

import (
	"CryptoClub/config"
	"CryptoClub/models"
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func StoreDataInMongo(price models.CryptoData, cfg config.Config) error {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	clientOptions.SetTLSConfig(&tls.Config{
		// Replace this with proper TLS configuration in production
		InsecureSkipVerify: true,
	})

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return fmt.Errorf("failed to create new client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to client: %w", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(cfg.DBName).Collection(cfg.CollectionName)

	document := bson.M{
		"timestamp": time.Now(),
		"currency":  price.Currency,
		"price":     price.Price,
	}

	_, err = collection.InsertOne(ctx, document)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}

	fmt.Printf("Stored - Currency: %s, Price: %f\n", price.Currency, price.Price)
	return nil
}
