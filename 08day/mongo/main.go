package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// //设置连接客户端配置
	// clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	// //连接到mongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //检查连接
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("connect successed")

	credential := options.Credential{
		Username: "hsiaocz",
		Password: "shaw123",
	}
	clientOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017").SetAuth(credential)
	// 还可以使用这种
	// clientOpts := options.Client().ApplyURI("mongodb://username:password@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connect successed")
}
