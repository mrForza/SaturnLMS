package dal

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *sqlx.DB

func InitDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(nil, nil)
	if err != nil {
		log.Fatal("Ошибка подключения к MongoDB")
	}

	collection := client.Database("filedb").Collection("files")
	log.Println("Подключено к MongoDB!")
	return nil
}
