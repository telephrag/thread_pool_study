package jobwithstate

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
)

type Job struct {
	State int64
	Mu    sync.Mutex
}

func New() *Job {
	return &Job{
		State: math.MaxInt64,
	}
}

func (j *Job) Do() {

	val := rand.Int63() % 10000
	ng := int64(runtime.NumGoroutine())

	work := md5.Sum([]byte(fmt.Sprint(ng + val))) // h(h(m) || m)
	work = md5.Sum([]byte(fmt.Sprint(work, ng+val)))

	val = int64(binary.BigEndian.Uint16(work[:2])) % 10000

	j.Mu.Lock()
	defer j.Mu.Unlock()
	if val < j.State {
		//fmt.Printf("swap: %d, %d\n", val, j.State)
		j.State = val
	}
}
