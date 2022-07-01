package dbs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"github.com/MKA-Nigeria/mkanmedia-go/config"
	mgo "go.mongodb.org/mongo-driver/mongo"
)

//ConnectMongodb returns a connection to a mongodb instance through a connection string in the configuration file
func ConnectMongodb() *mgo.Client {
	url := config.Env.MONGO_URL

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mgo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Printf("db connection error: %s", err.Error())
	}

	dbs, err := client.ListDatabases(context.Background(), nil)
	log.Printf("db connection: %s", dbs)
	return client
}
