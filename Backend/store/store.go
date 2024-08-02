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

// MongoClient is a singleton MongoDB client
var mongoClient *mongo.Client

// InitializeMongoClient initializes the MongoDB client
func InitializeMongoClient(cfg *config.Config) error {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	clientOptions.SetTLSConfig(&tls.Config{
		InsecureSkipVerify: true, // Adjust as needed for production
	})

	var err error
	mongoClient, err = mongo.NewClient(clientOptions)
	if err != nil {
		return fmt.Errorf("failed to create new client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = mongoClient.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to client: %w", err)
	}

	return nil
}

// StoreDataInMongo stores a CryptoData instance in MongoDB
func StoreDataInMongo(price models.CryptoData, cfg *config.Config) error {
	if mongoClient == nil {
		return fmt.Errorf("MongoClient is not initialized")
	}

	collection := mongoClient.Database(cfg.DBName).Collection(cfg.CollectionName)

	document := bson.M{
		"timestamp": time.Now(),
		"currency":  price.Currency,
		"price":     price.Price,
	}

	_, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}

	fmt.Printf("Stored - Currency: %s, Price: %f\n", price.Currency, price.Price)
	return nil
}

// FetchRecentData retrieves the most recent 20 entries from MongoDB
func FetchRecentData(cfg *config.Config) ([]models.CryptoData, error) {
	if mongoClient == nil {
		return nil, fmt.Errorf("MongoClient is not initialized")
	}

	collection := mongoClient.Database(cfg.DBName).Collection(cfg.CollectionName)

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"timestamp", -1}})
	findOptions.SetLimit(20)

	cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find documents: %w", err)
	}
	defer cursor.Close(context.Background())

	var results []models.CryptoData
	for cursor.Next(context.Background()) {
		var result models.CryptoData
		err := cursor.Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("failed to decode document: %w", err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return results, nil
}
