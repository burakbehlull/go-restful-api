package utils

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Db(){

	mongoUri := os.Getenv("MONGO_URI")
	
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("MongoDB bağlantı hatası:", err)
		return
	}
	if client == nil {
		fmt.Println("Client boş, bağlantı yok")
	} else {
		fmt.Println("Client dolu, bağlantı var")
	}

}