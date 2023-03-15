package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client   *mongo.Client                                               //handle client lifecycle
	UserColl *mongo.Collection                                           //export Collection
	Ctx, _   = context.WithTimeout(context.Background(), 10*time.Second) //export context
)

func InitializeDb() {
	//env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("mongo_uri")
	dbName := os.Getenv("db_name")
	collName := os.Getenv("user_coll")
	if len(uri) == 0 || len(dbName) == 0 || len(collName) == 0 {
		panic("something is missing")
	}

	//db
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	Client = client

	err = Client.Connect(Ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(Ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db := Client.Database(dbName)
	UserColl = db.Collection(collName)
}
