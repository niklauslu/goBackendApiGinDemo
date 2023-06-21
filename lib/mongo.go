package lib

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func MongoClientGet() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, errors.New("env error")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func MongoDatabaseGet() (*mongo.Database, error) {
	dbName := os.Getenv("MONGODB_DATABASE")

	if dbName == "" {
		return nil, errors.New("env error")
	}

	client, err := MongoClientGet()
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	return db, nil
}

func MongoCollection(collName string) (*mongo.Collection, error) {
	db, err := MongoDatabaseGet()
	if err != nil {
		return nil, err
	}
	coll := db.Collection(collName)
	return coll, nil
}

func MongoDBWithTransaction(fn func(ctx mongo.SessionContext, client *mongo.Client) (interface{}, error)) error {
	wc := writeconcern.New(writeconcern.WMajority())
	txnOptions := options.Transaction().SetWriteConcern(wc)

	client, err := MongoClientGet()
	if err != nil {
		return err
	}
	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.TODO())

	_, err = session.WithTransaction(context.TODO(), func(sessCtx mongo.SessionContext) (interface{}, error) {
		result, err := fn(sessCtx, client)
		return result, err
	}, txnOptions)

	return err
}
