package workerpool

import "sync"

type WorkerPool struct {
	RunningJobs chan rune
}

func New(threadCount uint64) *WorkerPool { // TODO: Handle error when threadCount = 0
	return &WorkerPool{
		RunningJobs: make(chan rune, threadCount-1),
	}
}

func (wp *WorkerPool) DoWork(work func(*sync.Mutex), wg *sync.WaitGroup, mu *sync.Mutex) { // TODO: Add case check wether worker pool has finished
	wp.RunningJobs <- 1
	go func() {
		work(mu)
		<-wp.RunningJobs // free the slot for another job
		wg.Done()        // TODO: Move WaitGroup inside the WorkerPool struct
	}()
}
