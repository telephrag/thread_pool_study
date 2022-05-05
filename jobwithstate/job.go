package jobwithstate

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
)

type Job struct {
	State int64
}

func New() *Job {
	return &Job{
		State: math.MaxInt64,
	}
}

func (j *Job) Do(mu *sync.Mutex) {

	val := rand.Int63() % 10000
	ng := int64(runtime.NumGoroutine())

	work := md5.Sum([]byte(fmt.Sprint(ng + val))) // h(h(m) || m)
	work = md5.Sum([]byte(fmt.Sprint(work, ng+val)))

	val = int64(binary.BigEndian.Uint16(work[:2])) % 10000

	swapped := atomic.CompareAndSwapInt64(&(j.State), j.State, val)
	if swapped {
		fmt.Printf("new state: %d\n", j.State)
	}
}
