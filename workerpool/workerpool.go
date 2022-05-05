package workerpool

import (
	"sync"
)

type WorkerPool struct {
	RunningJobs chan rune
	Wg          sync.WaitGroup
}

func New(threadCount uint64) *WorkerPool { // TODO: Handle error when threadCount = 0
	return &WorkerPool{
		RunningJobs: make(chan rune, threadCount-1),
	}
}

func (wp *WorkerPool) DoWork(work func()) { // TODO: Asser weather worker pool has finished
	wp.RunningJobs <- 1
	wp.Wg.Add(1)
	go func() {
		work()
		<-wp.RunningJobs // free the slot for another job
		wp.Wg.Done()
	}()
}

func (wp *WorkerPool) Await() { // TODO: Add wait timeout. Make it dynamically wait for tasks.
	if len(wp.RunningJobs) > 0 {
		wp.Wg.Wait()
	}
}

func (wp *WorkerPool) Resume() {
	wp.Wg.Done() // TODO
}
