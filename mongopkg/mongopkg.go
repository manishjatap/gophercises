package mongopkg

//package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoOps : Interface to perform mongo db operations
type MongoOps interface {
	Insert()
	Delete()
	Update()
	Get() []MongoOps
}

//Task : Structure of task type
type Task struct {
	Name           string
	CreateDate     string
	Status         string
	CompletionDate string
}

//getConnection : Return the connection object to MongoDB [Not Exposed]
func getConnection() *mongo.Client {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Connected to MongoDB!")

	return client
}

//getConnection : Return the connection object to MongoDB [Not Exposed]
func disconnect(client *mongo.Client) {

	// Disconnect from MongoDB
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Disonnected from MongoDB!")
}

//Insert : Insert a document into a collection
func (t Task) Insert() {

	client := getConnection()

	t.Status = "incomplete"
	t.CreateDate = time.Now().String()

	collection := client.Database("taskdb").Collection("tasks")

	_, err := collection.InsertOne(context.TODO(), t)

	if err != nil {
		log.Fatal(err)
	}

	disconnect(client)
}

//Get : Fetch all  the task from 'Tasks' collection
func (t Task) Get() []MongoOps {

	//Get the connection
	client := getConnection()

	//Prepare collection type
	collection := client.Database("taskdb").Collection("tasks")

	var taskList []MongoOps

	filter := bson.M{
		"status": bson.M{
			"$eq": "incomplete",
		},
	}

	cur, err := collection.Find(context.TODO(), filter, options.Find())

	if err != nil {
		log.Fatal(err) //Error handling
	}

	for cur.Next(context.TODO()) {

		var myTask Task

		err := cur.Decode(&myTask)

		if err != nil {
			log.Fatal(err) //Error handling
		}

		taskList = append(taskList, myTask)
	}

	cur.Close(context.TODO())

	disconnect(client)

	return taskList
}

//Delete : Remove document from collection
func (t Task) Delete() {

	client := getConnection()

	collection := client.Database("taskdb").Collection("tasks")

	_, err := collection.DeleteOne(context.TODO(), bson.M{"name": t.Name})

	if err != nil {
		log.Fatal(err)
	}

	disconnect(client)
}

//Update : Update the document from the collection
func (t Task) Update() {

	client := getConnection()

	collection := client.Database("taskdb").Collection("tasks")

	filter := bson.M{
		"name": bson.M{
			"$eq": t.Name,
		},
	}

	update := bson.M{
		"$set": bson.M{
			"status":         "completed",
			"completiondate": time.Now().String(),
		},
	}

	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	disconnect(client)
}

// func main() {
// 	fmt.Println("Mongo")

// 	t := new(Task)
// 	//t.Name = "Go to school"

// 	fmt.Println(t.Get())
// 	//GetConnection()
// 	//Insert("Prepare a tea!")
// 	//Insert("Prepare a coffee!")cd ..
// 	//Insert("Prepare a rice!")
// 	//Get()
// 	//Delete("Prepare a rice!")
// 	//Update("Prepare a conffee!", false)

// }
