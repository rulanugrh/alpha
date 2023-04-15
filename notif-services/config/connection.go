package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotifDB struct {
	Name     string
	Password string
	Host     string
}

var notifDB *NotifDB

func MngoConnection() *mongo.Client {
	conf := GetConfig()

	login := fmt.Sprintf("mongodb+src://%s:%s@%s/?retryWrites=true&w=majority", conf.Name, conf.Password, conf.Host)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(login).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func MngoColletion(coll string, client *mongo.Client) *mongo.Collection {
	return client.Database("alpha").Collection(coll)
}

func GetConfig() *NotifDB {
	if notifDB == nil {
		notifDB = initConfig()
	}
	return notifDB
}

func initConfig() *NotifDB {
	conf := NotifDB{}
	err := godotenv.Load()
	if err != nil {
		conf.Host = "localhost"
		conf.Password = "admin"
		conf.Name = "admin"
		return &conf
	}

	conf.Host = os.Getenv("MONGODB_HOST")
	conf.Password = os.Getenv("MONGODB_PASS")
	conf.Name = os.Getenv("MONGODB_NAME")

	return &conf
}
