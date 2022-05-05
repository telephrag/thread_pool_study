package main

import (
	"fmt"
	"sync"
	"thread_pool_study/config"
	"thread_pool_study/jobwithstate"
	"thread_pool_study/workerpool"
	"time"

	"github.com/zenthangplus/goccm"
)

func plane() int64 {
	j := jobwithstate.New()
	var wg sync.WaitGroup

	wg.Add(int(config.ThreadCount) * config.Iterations)
	for i := 0; i < int(config.ThreadCount)*config.Iterations; i++ {
		go func() {
			j.Do()
			wg.Done()
		}()
	}
	wg.Wait()

	return j.State
}

func workerPool() int64 {
	wp := workerpool.New(config.ThreadCount * 2)
	j := jobwithstate.New()

	for i := 0; i < int(config.ThreadCount)*config.Iterations; i++ {
		wp.DoWork(j.Do)
	}
	wp.Await()

	return j.State
}

func goccMan() int64 {
	ccm := goccm.New(int(config.ThreadCount))
	j := jobwithstate.New()

	for i := 0; i < int(config.ThreadCount)*config.Iterations; i++ {
		ccm.Wait()
		go func() {
			j.Do()
			ccm.Done()
		}()
	}
	ccm.WaitAllDone()

	return j.State
}

func measureExecTime(f func() int64) {
	start := time.Now()
	res := f()
	elapsed := time.Since(start)
	fmt.Printf("finished in %d us\n", elapsed.Microseconds())
	fmt.Printf("result: %d\n", res)
}

func main() {
	fmt.Println("\nplane")
	measureExecTime(plane)

	fmt.Println("\nworkerpool")
	measureExecTime(workerPool)

	fmt.Println("\ngoccm")
	measureExecTime(goccMan)
}
