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
	State int
}

func New() *Job {
	return &Job{
		State: math.MaxInt,
	}
}

func (j *Job) Do() {

	val := rand.Int() % 10000
	ng := runtime.NumGoroutine()
	work := md5.Sum([]byte(fmt.Sprint(ng + val))) // h(h(m) || m)
	work = md5.Sum([]byte(fmt.Sprint(work, ng+val)))
	val = int(binary.BigEndian.Uint16(work[:2])) % 10000

	// fmt.Printf("goroutines exist: %d; work: %v\n", ng, work)
	// is time.Sleep() bad representation of work being done?
	// consider using hash-function calculation instead -> apparently func requesres implementation
	// time.Sleep(time.Millisecond * 2000)

	//fmt.Printf("val: %d\n", val)
	if val < j.State {
		var mu sync.Mutex
		mu.Lock()
		defer mu.Unlock()
		j.State = val
	}
}
