package queue

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const envDatabase = "MONGO_DATABASE"
const envCollection = "MONGO_QUEUE"

type MongoDB struct {
	col *mongo.Collection
}

func NewMongoDB(col *mongo.Collection) *MongoDB {
	return &MongoDB{col}
}

func NewMongoURI(URI string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(URI)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		client.
			Database(os.Getenv(envDatabase)).
			Collection(os.Getenv(envCollection))}, nil
}

func (MongoDB MongoDB) Restore() ([]*Job, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := MongoDB.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	jobs := []*Job{}
	for cur.Next(ctx) {
		j := &Job{}
		err := cur.Decode(j)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, j)
	}

	return jobs, nil
}

func (MongoDB MongoDB) Store(jobs []*Job) error {
	ctx := context.Background()
	for _, j := range jobs {
		_, err := MongoDB.col.InsertOne(ctx, j)
		if err != nil {
			return err
		}
	}
	return nil
}
