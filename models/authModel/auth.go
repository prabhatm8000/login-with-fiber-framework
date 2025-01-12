package authModel

import (
	"context"
	"fmt"

	"example.com/login/configs/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

var CollectionName = "users"
var Collection *mongo.Collection

func (u *User) Create() (primitive.ObjectID, error) {
	if Collection == nil {
		Collection = mongodb.GetCollection(CollectionName)
	}

	// Check if email is unique
	filter := bson.D{{Key: "email", Value: u.Email}}
	count, err := Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return primitive.NilObjectID, err
	}
	if count > 0 {
		return primitive.NilObjectID, fmt.Errorf("email already exists")
	}

	// Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return primitive.NilObjectID, err
	}

	u.Password = string(hash)

	result, err := Collection.InsertOne(context.Background(), u)
	if err != nil {
		return primitive.NilObjectID, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid, nil
	}

	return primitive.NilObjectID, mongo.ErrNilDocument
}

func (u *User) FindByID(id string) error {
	if Collection == nil {
		Collection = mongodb.GetCollection(CollectionName)
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "_id", Value: oid}}

	var user User
	err = Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return err
	}

	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.Password = ""
	return nil
}

func (u *User) FindByEmail() error {
	if Collection == nil {
		Collection = mongodb.GetCollection(CollectionName)
	}
	filter := bson.D{{Key: "email", Value: u.Email}}
	err := Collection.FindOne(context.Background(), filter).Decode(&u)
	return err
}

func (u *User) MatchPassword(password string) (bool, error) {
	if Collection == nil {
		Collection = mongodb.GetCollection(CollectionName)
	}
	filter := bson.D{{Key: "email", Value: u.Email}}
	err := Collection.FindOne(context.Background(), filter).Decode(&u)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
