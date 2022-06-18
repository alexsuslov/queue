package main

import (
	"fmt"
	"github/alexsuslov/queue"
	"time"
)

func main() {
	fmt.Println("start queue function test")
	db, err := queue.NewBoltDB("test.db")
	if err != nil {
		panic(err)
	}
	//	new queue
	q, err := queue.New(db)
	if err != nil {
		panic(err)
	}

	//	new job
	t := time.NewTicker(1 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println("new job")
		<-t.C
		j := queue.NewJob(i)
		q.Append(j)
	}
	q.Jobs[3].Cancel = true
	time.Sleep(5 * time.Second)
	fmt.Printf("q=%v", q)

}
