package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Account struct {
	ID      string  `bson:"_id"`
	Balance float64 `bson:"balance"`
}

func main() {
	// Thiết lập thông tin kết nối MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://172.31.96.1:10001,172.31.96.1:10002,172.31.96.1:10003")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Chọn cơ sở dữ liệu và collection cần thao tác
	db := client.Database("mydb")
	collection := db.Collection("mycollection")

	// Tạo hai tài khoản trên hai shard khác nhau
	account1 := Account{ID: "X", Balance: 100.0}
	account2 := Account{ID: "Y", Balance: 200.0}

	// Ghi tài khoản 1 lên shard 1
	_, err = collection.InsertOne(context.Background(), account1)
	if err != nil {
		log.Fatal(err)
	}

	// Ghi tài khoản 2 lên shard 2
	_, err = collection.InsertOne(context.Background(), account2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Đã tạo hai tài khoản trên hai shard khác nhau!")
}
