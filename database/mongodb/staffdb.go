package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/waseemofficial/GraphQL-Redis-MongoDB-Go/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) SaveStaff(input *model.NewStaff) *model.Staff {
	collection := db.client.Database("School").Collection("Staff")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Staff{
		ID:           res.InsertedID.(primitive.ObjectID).Hex(),
		StaffName:    input.StaffName,
		Position:     input.Position,
		ClassAlloted: []*model.Classes{},
	}
}

func (db *DB) FindSaff(ID string) *model.Staff {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("School").Collection("Staff")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})

	staff := model.Staff{}
	res.Decode(&staff)
	return &staff
}
