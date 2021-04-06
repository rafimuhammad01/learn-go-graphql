package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type DB struct {
	client *mongo.Client
	Database *mongo.Database
}

func Connect() *DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	URI := "mongodb+srv://" +os.Getenv("DB_USERNAME") +":" + os.Getenv("PASSWORD") + "@" +os.Getenv("cluseter_address")+"/" + os.Getenv("db_name")+ "?retryWrites=true&w=majority"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil { log.Fatal(err) }

	return &DB{
		client: client,
		Database:     client.Database(os.Getenv("DB_NAME")),
	}
}



