package main

import (
	"fmt"
	"thread_pool_study/config"
	"thread_pool_study/jobwithstate"
	"thread_pool_study/workerpool"
	"time"

	"github.com/zenthangplus/goccm"
)

func plane(done chan int) int {
	j := jobwithstate.New()

	for i := 0; i < int(config.ThreadCount)*100; i++ {
		go j.Do()
	}

	done <- 1

	return j.State
}

func workerPool(done chan int) int {
	wp := workerpool.New(config.ThreadCount)
	j := jobwithstate.New()

	for i := 0; i < int(config.ThreadCount)*100; i++ {
		wp.DoWork(j.Do)
	}

	done <- 1

	return j.State
}

func goccMan() int {
	ccm := goccm.New(int(config.ThreadCount))
	j := jobwithstate.New()

	for i := 0; i < int(config.ThreadCount*100); i++ {
		ccm.Wait()
		go func(index int) {
			j.Do()
			ccm.Done()
		}(i)
	}
	ccm.WaitAllDone()

	return j.State
}

func main() {

	done := make(chan int, 1)
	start := time.Now()
	res := workerPool(done)
	<-done
	elapsed := time.Since(start)
	fmt.Println("\nthreadpool")
	fmt.Printf("finished in %d micro-seconds\n", elapsed.Microseconds())
	fmt.Printf("result: %d\n", res)

	start = time.Now()
	res = plane(done)
	<-done
	elapsed = time.Since(start)
	fmt.Println("\nplane")
	fmt.Printf("finished in %d micro-seconds\n", elapsed.Microseconds())
	fmt.Printf("result: %d\n", res)

	start = time.Now()
	res = goccMan()
	elapsed = time.Since(start)
	fmt.Println("\ngoccm")
	fmt.Printf("completed in %dmcs\n", elapsed.Microseconds())
	fmt.Printf("result %d\n", res)
}
