package db

import (
	"context"
	"errors"
	"time"

	model "runner/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cursor *mongo.Cursor
var err error

func GetUser(email string) primitive.M {
	var user bson.M

	err = UserColl.FindOne(context.TODO(), bson.D{
		{Key: "email", Value: email},
	}).Decode(&user)

	if err != nil {
		return nil
	}
	return user
}

func GetAllUser() []primitive.M {
	var users []bson.M

	cursor, err = UserColl.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil
	}
	return users
}

func Insert(body *model.User) error {
	user := GetUser(body.Email)
	if user != nil {
		return errors.New("user exist")
	}

	_, err = UserColl.InsertOne(context.TODO(), model.User{
		Email: body.Email, Password: body.Password, CreatedAt: time.Now(), UpdatedAt: time.Now(),
	})
	if err != nil {
		return errors.New("user not addes")
	}
	return nil
}

func BulkAdd() error {
	_, err = UserColl.InsertMany(context.TODO(), []interface{}{
		model.User{Email: "test1@gmail.com", Password: "hi"},
		model.User{Email: "test2@gmail.com", Password: "hi"},
		model.User{Email: "test3@gmail.com", Password: "hi"},
	})
	if err != nil {
		return errors.New("users not added")
	}
	return nil
}

func UpdateOne(email string, kv map[string]any) error {
	user := GetUser(email)
	if user == nil {
		return errors.New("user not found")
	}
	var key string
	var val any

	for k, v := range kv {
		key = k
		val = v
	}

	_, err = UserColl.UpdateOne(context.TODO(),
		bson.D{
			{Key: "email", Value: email},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{
					Key:   key,
					Value: val,
				},
				{
					Key:   "updated_at",
					Value: time.Now(),
				},
			}},
		})
	if err != nil {
		return errors.New("not updated")
	}
	return nil
}

func Delete(email string) error {
	user := GetUser(email)
	if user == nil {
		return errors.New("not found user")
	}
	_, err = UserColl.DeleteOne(context.TODO(),
		bson.D{
			{Key: "email", Value: email},
		},
	)
	if err != nil {
		return errors.New("delte error")
	}
	return nil
}

func DeleteAll() error {
	_, err := UserColl.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		return errors.New("not implemented")
	}
	return nil
}
