package mongopkg

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoOps : Interface to perform mongo db operations
type MongoOps interface {
	Insert()
	Delete()
	Update()
	Get() []Task
}

//Task : Structure of task type
type Task struct {
	Name           string
	CreateDate     string
	Status         string
	CompletionDate string
}

//getSession : Return the connection object to MongoDB
var getSession = func() (*mgo.Session, error) {
	return mgo.Dial("mongodb://localhost:27017")
}

//closeSession : Close the MongoDB session
var closeSession = func(s *mgo.Session) {
	// Disconnect from MongoDB
	s.Close()
}

//insertDoc : Insert document into MongoDB collection
var insertDoc = func(collection *mgo.Collection, t Task) error {
	return collection.Insert(t)
}

//removeDoc : Remove document from MongoDB collection
var removeDoc = func(collection *mgo.Collection, filter interface{}) error {
	return collection.Remove(filter)
}

//updateDoc : Update document from MongoDB collection
var updateDoc = func(collection *mgo.Collection, filter interface{}, update interface{}) error {
	return collection.Update(filter, update)
}

var getDoc = func(collection *mgo.Collection, filter interface{}, tlist *[]Task) error {
	iter := collection.Find(filter).Iter()
	return iter.All(tlist)
}

//Insert : Insert a document into a collection
func (t Task) Insert() error {

	session, err := getSession()

	if err != nil {
		return err
	}

	defer closeSession(session)

	//Prepare collection type
	collection := session.DB("taskdb").C("tasks")

	t.Status = "incomplete"
	t.CreateDate = time.Now().String()

	insertErr := insertDoc(collection, t)

	if insertErr != nil {
		return insertErr
	}

	return nil
}

//Get : Fetch all  the task from 'Tasks' collection
func (t Task) Get() ([]Task, error) {

	session, err := getSession()

	if err != nil {
		return nil, err
	}

	defer closeSession(session)

	//Prepare collection type
	collection := session.DB("taskdb").C("tasks")

	filter := bson.M{
		"status": bson.M{
			"$eq": "incomplete",
		},
	}

	var taskList []Task

	fetchErr := getDoc(collection, filter, &taskList)

	if fetchErr != nil {
		return taskList, fetchErr
	}

	return taskList, nil
}

//Delete : Remove document from collection
func (t Task) Delete() error {

	session, err := getSession()

	if err != nil {
		return err
	}

	defer closeSession(session)

	//Prepare collection type
	collection := session.DB("taskdb").C("tasks")

	removeErr := removeDoc(collection, bson.M{"name": t.Name})

	if removeErr != nil {
		return removeErr
	}

	return nil
}

//Update : Update the document from the collection
func (t Task) Update() error {

	session, err := getSession()

	if err != nil {
		return err
	}

	defer closeSession(session)

	//Prepare collection type
	collection := session.DB("taskdb").C("tasks")

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

	updateErr := updateDoc(collection, filter, update)

	if updateErr != nil {
		return updateErr
	}

	return nil
}
