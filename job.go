package queue

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Job struct {
	ID        primitive.ObjectID `json:"_id",bson:"_id"`
	CreatedON time.Time          `json:"created_on"`
	Cancel    bool               `json:"cancel"`
	Value     interface{}        `json:"value"`
}

func NewJob(v interface{}) *Job {
	return &Job{
		primitive.NewObjectID(),
		time.Now(),
		false,
		v,
	}

}

func (Job Job) Bytes() []byte {
	data, _ := json.Marshal(Job)
	return data
}

func Bytes2Job(data []byte) (*Job, error) {
	job := &Job{}
	return job, json.Unmarshal(data, job)
}
