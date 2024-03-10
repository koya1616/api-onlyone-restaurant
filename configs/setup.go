package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("restaurant").Collection(collectionName)
	return collection
}

var neonPostgresDb *gorm.DB

func GetNeonDB() *gorm.DB {
	if neonPostgresDb == nil {
		InitNeonDB()
	}
	return neonPostgresDb
}

func InitNeonDB() {
	var err error
	dsn := EnvNeonDBString()
	neonPostgresDb, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
	}
}
