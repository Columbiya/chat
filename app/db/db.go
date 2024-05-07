package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MessagesColl *mongo.Collection

func Connect() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	db := client.Database("zb-site")

	EnsureCreated(db)

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("zb-site").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func EnsureCreated(db *mongo.Database) {
	coll := db.Collection("messages")

	if coll == nil {
		if err := db.CreateCollection(context.Background(), "messages"); err != nil {
			log.Fatal(err)
			panic(err)
		}

		MessagesColl = db.Collection("messages")
		return
	}

	MessagesColl = coll
}
