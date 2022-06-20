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
	Retry     int                `json:"retry"`
	Value     interface{}        `json:"value"`
	Err       error              `json:"-"`
}

func NewJob(v interface{}) *Job {
	return &Job{
		primitive.NewObjectID(),
		time.Now(),
		false,
		0,
		v,
		nil,
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
