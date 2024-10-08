package utils

import (
	"fmt"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db(){
	clientOptions := options.Client().ApplyURI("MONGODB URL HERE")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("MongoDB bağlantı hatası:", err)
		return
	}

}