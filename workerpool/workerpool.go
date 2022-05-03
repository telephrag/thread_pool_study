package workerpool

type WorkerPool struct {
	RunningJobs chan rune
}

func New(threadCount uint64) *WorkerPool { // TODO: Handle error when threadCount = 0
	return &WorkerPool{
		RunningJobs: make(chan rune, threadCount-1),
	}
}

func (wp *WorkerPool) DoWork(work func()) { // TODO: Add case check wether worker pool has finished
	wp.RunningJobs <- 1
	go func() {
		work()
		<-wp.RunningJobs // free the slot for another job
	}()
}
