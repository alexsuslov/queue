package queue

import (
	"golang.org/x/exp/slices"
	"sync"
	"time"
)

const bucketName = "jobs"
const backupTimer = 1 * time.Second

type IDb interface {
	Restore() ([]*Job, error)
	Store([]*Job) error
}

type Queue struct {
	DB   IDb
	Jobs []*Job
	sync.RWMutex
	Updated bool
}

func New(DB IDb) (queue *Queue, err error) {
	queue = &Queue{DB: DB}

	//	restore
	if DB != nil {
		jobs, err := DB.Restore()
		if err != nil {
			return nil, err
		}

		queue.Jobs = jobs
	}

	// create backup
	t := time.NewTicker(backupTimer)
	go func() {
		for {
			<-t.C
			go queue.Store()
		}
	}()

	return queue, nil
}

func (Queue Queue) ActiveJobs() (c int) {
	for _, job := range Queue.Jobs {
		if !job.Cancel {
			c++
		}
	}
	return
}

func (Queue *Queue) Append(job *Job) (idx int) {
	Queue.Lock()
	Queue.Jobs = append(Queue.Jobs, job)
	Queue.Updated = true
	Queue.Unlock()
	return len(Queue.Jobs)
}

func (Queue *Queue) Remove(idx int) *Queue {
	Queue.Lock()
	Queue.Jobs = slices.Delete(Queue.Jobs, idx, 1)
	Queue.Updated = true
	Queue.Unlock()
	return Queue
}

func (Queue Queue) Store() error {
	if Queue.DB == nil || !Queue.Updated {
		return nil
	}
	err := Queue.DB.Store(Queue.Jobs)
	if err != nil {
		return err
	}
	Queue.Lock()
	Queue.Updated = false
	Queue.Unlock()

	return nil
}
