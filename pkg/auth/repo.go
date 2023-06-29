package auth

import (
	"context"
	"errors"
	"sigmacoder/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository is an interfaces that defines the schema of
// the CRUD operations that can be performed on the User
// entity. The Implementation might be changed later in
// case we migrate away from gorm.
type Repository interface {
	Create(in InUser) (User, error)
	Read(id string) (User, error)
	Update(id string, upd map[string]interface{}) (User, error)
	Delete(string int) bool
	ReadByID(id string) (User, error)
	ReadByEmail(email string) (User, error)
	ReadByPhoneNumber(phone string) (User, error)
	ReadByUsernanme(username string) (User, error)
}

// Repo is the struct that Implements the Repository Interface.
// To Create a Repo, Use the NewRepo Function, it takes in a DB of type *gorm.DB
type Repo struct {
	db      *mongo.Collection
	context context.Context
}

// This function is used to fetch a user from the database with their email. It takes in an email
// string as a parameter and returns a User object and an error. It searches for a user in the database
// with the given email using the FindOne method of the MongoDB collection. If a user is found, it
// decodes the result into a User object and returns it. If no user is found, it returns an error
// indicating that the user was not found.
func (s *Repo) ReadByEmail(email string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, pkg.ErrUserNotFound
	}
	return user, nil
}

// This function is used to fetch a user from the database with their phone number. It takes in a phone
// number string as a parameter and returns a User object and an error. It searches for a user in the
// database with the given phone number using the FindOne method of the MongoDB collection. If a user
// is found, it decodes the result into a User object and returns it. If no user is found, it returns
// an error indicating that the user was not found.

func (s *Repo) ReadByPhoneNumber(phone string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, bson.M{"phonenumber": phone}).Decode(&user)
	if err != nil {
		return user, pkg.ErrUserNotFound
	}
	return user, nil
}

// This function is used to fetch a user from the database with their username. It takes in a username
// string as a parameter and returns a User object and an error. It searches for a user in the database
// with the given username using the FindOne method of the MongoDB collection. If a user is found, it
// decodes the result into a User object and returns it. If no user is found, it returns an error
// indicating that the user was not found.
func (s *Repo) ReadByUsernanme(username string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return user, errors.New("user not found with this email")
	}
	return user, nil
}

// This function is used to fetch a user from the database with their ID. It takes in an ID string as a
// parameter and returns a User object and an error. It converts the ID string to a MongoDB ObjectID
// using the `primitive.ObjectIDFromHex` function, and then searches for a user in the database with
// the given ID using the `FindOne` method of the MongoDB collection. If a user is found, it decodes
// the result into a User object and returns it. If no user is found, it returns an error indicating
// that the user was not found.
func (s *Repo) ReadByID(id string) (User, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	var user User
	err := s.db.FindOne(s.context, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// This function is creating a new user in the database. It takes an `InUser` object as input, which is
// a struct that contains the necessary information to create a new user. It converts this `InUser`
// object to a `User` object using the `ToUser()` method, and then inserts this `User` object into the
// MongoDB collection using the `InsertOne()` method. If there is an error during the insertion, it
// returns the error. Otherwise, it returns the newly created `User` object.
func (s *Repo) Create(in InUser) (User, error) {
	user := in.ToUser()
	_, err := s.db.InsertOne(s.context, user)
	if err != nil {
		return user, err
	}
	return user, nil

}

// `func (s *Repo) Read(id string) (User, error)` is a method of the `Repo` struct that implements the
// `Repository` interface. It takes an `id` of type `string` as input and returns a `User` object and
// an `error`. It searches for a user in the database with the given ID using the `FindOne` method of
// the MongoDB collection. If a user is found, it decodes the result into a `User` object and returns
// it. If no user is found, it returns an error indicating that the user was not found.
func (s *Repo) Read(id string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return user, errors.New("user not found with this id")
	}
	return user, nil
}

// This function is updating a user in the database. It takes in an ID string and a map of fields to
// update as input. It searches for a user in the database with the given ID using the
// `FindOneAndUpdate` method of the MongoDB collection, and updates the fields specified in the input
// map. If the update is successful, it returns the updated `User` object. If there is an error during
// the update, it returns the error.
func (s *Repo) Update(id string, upd map[string]interface{}) (User, error) {
	var u User
	if err := s.db.FindOneAndUpdate(s.context, bson.M{"_id": id}, upd).Decode(&u); err != nil {
		return u, err
	}
	return u, nil
}

// `func (s *Repo) Delete(id int) bool` is a method of the `Repo` struct that implements the
// `Repository` interface. It takes an `id` of type `int` as input and returns a `bool`.
func (s *Repo) Delete(id int) bool {
	delete, err := s.db.DeleteOne(s.context, bson.M{"_id": id})
	if err != nil {
		return false
	}
	return delete.DeletedCount == 1
}

// The function returns a new instance of a Repository interface implementation with a MongoDB database
// connection.
func NewRepo(db *mongo.Database) Repository {
	ctx := context.TODO()
	return &Repo{db: db.Collection("users"), context: ctx}
}
