package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Yashh56/todo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionSting = " MongoDB URL"
const dbName = "TodoCLI"
const colName = "Todos"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionSting)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName)
}

func AddTodo(todo model.Model) (string, error) {
	todo.CreatedAt = time.Now()
	todo.Done = false
	added, err := collection.InsertOne(context.Background(), todo)

	if err != nil {
		log.Fatal(added, err)
	}
	return " Todo has been added", nil
}

func DeleteTodo(title string) (string, error) {
	filter := bson.M{"title": title}
	delCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
		return " ", nil
	}

	if delCount.DeletedCount == 0 {
		return "No todo found with the given title", nil
	}

	return "Todo has been delete", nil
}

func DeleteAllTodos() (string, error) {
	filter := bson.D{}
	result, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Deleted %d todos", result.DeletedCount), nil
}

func UpdateTodo(title string) (string, error) {
	filter := bson.M{"title": title}
	update := bson.M{"$set": bson.M{"done": true}}

	if title == "" {
		fmt.Println("Not Found")
	}

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
		fmt.Println(res)
	}

	return "Todo has been done ", nil

}

func GetPendingTodo() ([]model.Model, error) {
	var todos []model.Model
	filter := bson.M{"done": false}
	cursor, err := collection.Find(context.Background(), filter, options.Find())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo model.Model
		err := cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return todos, nil
}
func GetAllDoneTodos() ([]model.Model, error) {
	var todos []model.Model
	filter := bson.M{"done": true}
	cursor, err := collection.Find(context.Background(), filter, options.Find())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo model.Model
		err := cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return todos, nil
}
func GetAllTodos() ([]model.Model, error) {
	var todos []model.Model
	cursor, err := collection.Find(context.Background(), bson.M{}, options.Find())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo model.Model
		err := cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return todos, nil
}
