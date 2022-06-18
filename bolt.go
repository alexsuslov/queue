package queue

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
)

type BoltDB struct {
	db *bolt.DB
}

func NewBoltDB(filename string) (db *BoltDB, err error) {
	if filename == "" {
		return nil, fmt.Errorf("filename nil")
	}
	DB, err := bolt.Open(filename, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &BoltDB{DB}, nil
}

func (BoltDB BoltDB) Restore() (jobs []*Job, err error) {
	jobs = []*Job{}
	return jobs, BoltDB.db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return nil
		}
		return b.ForEach(func(_, v []byte) error {
			job, err := Bytes2Job(v)
			if err != nil {
				return err
			}
			jobs = append(jobs, job)
			return nil
		})
	})
}

func (BoltDB BoltDB) Store(jobs []*Job) error {
	return BoltDB.db.Update(func(tx *bolt.Tx) error {
		tx.DeleteBucket([]byte(bucketName))
		b, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return err
		}
		for _, job := range jobs {
			if job.Cancel {
				continue
			}
			err = b.Put([]byte(job.ID.Hex()), job.Bytes())
			if err != nil {
				return err
			}
		}
		return nil
	})
}
