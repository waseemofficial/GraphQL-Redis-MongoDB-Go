package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/waseemofficial/GraphQL-Redis-MongoDB-Go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	return &DB{
		client: client,
	}
}
func (db *DB) SaveStudent(input *model.NewStudent) *model.Student {
	collection := db.client.Database("School").Collection("Students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Student{
		ID:          res.InsertedID.(primitive.ObjectID).Hex(),
		StudentName: input.StudentName,
		Gender:      input.Gender,
		PhoneNo:     input.PhoneNo,
		RollNo:      input.RollNo,
		FeePayed:    false,
		Class:       &model.Classes{},
	}
}

func (db *DB) FindStudent(ID string) *model.Student {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("School").Collection("Students")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	fmt.Println(res, ctx)
	student := model.Student{}
	res.Decode(&student)
	fmt.Println(bson.M{"_id": ObjectID})
	fmt.Println(&student)
	return &student
}

func (db *DB) AllStudents() []*model.Student {
	collection := db.client.Database("School").Collection("Students")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var students []*model.Student
	for cur.Next(ctx) {
		var stud *model.Student
		err := cur.Decode(&stud)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, stud)
	}

	return students
}
