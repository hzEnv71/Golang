package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Student struct {
	ID   int    `bson:"_id"`
	Age  int    `bson:"age"`
	Name string `bson:"name"`
	Addr string `bson:"address"`
}

var client *mongo.Client
var coll *mongo.Collection

func main() {
	// 1.连接mongodb
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017").
			SetConnectTimeout(5*time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		fmt.Println(err)
		return
	}

	// 由于都是对同一个集合进行操作，所以在初始化mongodb时就选择了集合，防止后续出现大量重复代码
	coll = client.Database("db").Collection("stu")
	InsertStudent(&Student{1, 20, "张飞", "河南"})
	//DeleteStudentByID(1)
	QueryStudentByID(1)
	UpdateStudent(&Student{1, 27, "张飞", "郑州"})
	QueryStudentByID(1)
}
func InsertStudent(p *Student) error {
	// 插入文档
	if _, err := coll.InsertOne(context.Background(), p); err != nil {
		return err
	}
	fmt.Println("Document inserted successfully!")
	return nil
}

func InsertStudents(p []interface{}) error {

	// 插入多条文档
	if _, err := coll.InsertMany(context.Background(), p); err != nil {
		return err
	}
	fmt.Println("Document inserted successfully!")
	return nil

}
func DeleteStudentByID(p int) error {

	// 过滤条件
	fil := bson.M{"_id": p}

	// 删除文档
	if _, err := coll.DeleteOne(context.Background(), fil); err != nil {
		return err
	}
	fmt.Println("Document Delete successfully!")
	return nil

}
func UpdateStudent(p *Student) error {

	// 根据条件筛选要更新的文档
	filter := bson.M{"_id": p.ID}
	update := bson.M{"$set": p}

	// 更新文档
	if _, err := coll.UpdateOne(context.Background(), filter, update); err != nil {
		return err
	}
	fmt.Println("Document Update successfully!")
	return nil

}
func QueryStudentByID(p int) error {
	filter := bson.M{"_id": p} // 根据条件筛选要更新的文档

	var s Student
	// 插入文档
	err := coll.FindOne(context.Background(), filter).Decode(&s)
	if err != nil {
		return err
	}

	fmt.Println("Document Find successfully!")
	fmt.Printf("Document Find: %+v", s)

	return nil

}

func GetClient() *mongo.Client {
	return client
}

func Close() {
	_ = client.Disconnect(context.Background())
}
