package allquestions

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	ReadAllQuestion() ([]AllQuestion, error)
	ReadByID(id string) (AllQuestion, error)
}

type Repo struct {
	db      *mongo.Collection
	context context.Context
}

// The `ReadByID` function is a method of the `Repo` struct that implements the `Repository` interface.
// It is used to retrieve a single question from the MongoDB collection based on its ID.
func (s *Repo) ReadByID(id string) (AllQuestion, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	var user AllQuestion
	err := s.db.FindOne(s.context, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// The `ReadAllQuestion` function is a method of the `Repo` struct that implements the `Repository`
// interface. It is used to retrieve all the questions from the MongoDB collection.
func (s *Repo) ReadAllQuestion() ([]AllQuestion, error) {
	var allquestions []AllQuestion
	cursor, err := s.db.Find(s.context, bson.M{})
	if err != nil {
		return allquestions, err
	}
	for cursor.Next(s.context) {
		var allquestion AllQuestion
		cursor.Decode(&allquestion)
		allquestions = append(allquestions, allquestion)
	}
	return allquestions, nil
}

func NewRepo(db *mongo.Database) Repository {
	ctx := context.TODO()
	return &Repo{db: db.Collection("AllQuestion"), context: ctx}
}
