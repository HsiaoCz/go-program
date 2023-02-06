package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbname   = "test"
	collName = "article"
)

func main() {
	// 连接
	client, err := InitMongo()
	if err != nil {
		log.Fatal(err)
	}

	// 获取集合名称
	colltionNames, err := GetCollectionName(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("collName:", colltionNames)

	// 插入单条数据
	insertOne, err := InserOne(client, dbname, collName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("id", insertOne.InsertedID)

	// 插入多条数据
	insertMany, err := InsertMany(client, dbname, collName)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("insert Many:",insertMany.InsertedIDs)
}

// 通过username 和password 连接mongo
func InitMongo() (*mongo.Client, error) {

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
	return client, err
}

// 获取集合名称 返回的是一个[]string
func GetCollectionName(client *mongo.Client) ([]string, error) {
	// bson 主要用来写查询的筛选条件filter，构造文件记录以及接受查询解码的值，也就是go和mongo之间的序列化
	// 一般会使用这三种数据结构
	// bson.D{}:对文档的有序描述,key-value以逗号分隔
	// bson.M{}:map结构，key-value以冒号分隔，无序，使用最方便
	// bson.A{}:数组结构，元素要求是有序的文档描述，也就是元素是bson.D{}类型

	// 获取db
	db := client.Database("test")
	collectionName, err := db.ListCollectionNames(context.TODO(), bson.M{})
	return collectionName, err
}

// 单条插入数据
func InserOne(client *mongo.Client, dbname string, collectionName string) (*mongo.InsertOneResult, error) {
	collection := client.Database("test").Collection("article")
	// InsertOne
	insertOneResult, err := collection.InsertOne(context.TODO(), bson.M{"name": "lisi", "gender": "nan", "level": 23})
	return insertOneResult, err
}

// 插入多条数据
func InsertMany(client *mongo.Client, dbname string, collectionName string) (*mongo.InsertManyResult, error) {
	// InsertMany
	docs := []interface{}{
		bson.M{"name": "5t5", "gender": "男", "level": 0},
		bson.M{"name": "奈奈米", "gender": "男", "level": 1},
	}
	// Ordered 设置为false表示其中一条插入失败不会影响其他文档的插入，默认为true，一条失败其他都不会被写入
	insertManyOpts := options.InsertMany().SetOrdered(false)
	collection := client.Database(dbname).Collection(collectionName)
	insertManyResult, err := collection.InsertMany(context.TODO(), docs, insertManyOpts)
	return insertManyResult, err
}
