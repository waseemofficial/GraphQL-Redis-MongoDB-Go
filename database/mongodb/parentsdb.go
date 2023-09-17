package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/waseemofficial/GraphQL-Redis-MongoDB-Go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) SaveParent(input *model.NewParent) *model.Parent {
	collection := db.client.Database("School").Collection("Parents")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Parent{
		ID:      res.InsertedID.(primitive.ObjectID).Hex(),
		Name:    input.Name,
		PhoneNo: input.PhoneNo,
		Job:     input.Job,
		Spous:   []*model.Parent{},
	}
}

func (db *DB) FindParent(ID string) *model.Parent {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("School").Collection("Parents")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})

	parent := model.Parent{}
	res.Decode(&parent)
	return &parent
}
