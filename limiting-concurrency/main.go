package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strings"
	"sync"
	"time"
)

func request(url string) (string, error) {
	// Wait up to 1 second to send a reply.
	time.Sleep(time.Duration(rand.IntN(1000)) * time.Millisecond)

	if strings.Contains(url, "fail") {
		return "", errors.New("request failed")
	}

	return url, nil
}

func main() {
	urls := []string{
		"1",
		"2",
		"fail",
		"3",
		"4",
		"5",
		"fail",
	}

	fmt.Println("unlimited:")
	noLimit(urls)

	fmt.Println("\nlimited:")
	limited(urls)

	fmt.Println("\npool:")
	jobs := make([]PoolJob, 0)
	for _, url := range urls {
		jobs = append(jobs, func() (string, error) { return request(url) })
	}
	pool := NewLimitedPool(2, jobs...)
	pool.Run()

	for {
		select {
		case result := <-pool.Results:
			fmt.Println("result:", result)
		case err := <-pool.Errors:
			fmt.Println(err.Error())
		}

		if pool.IsEmpty() == true {
			break
		}
	}
}

// noLimit is a basic implementation of utilizing concurrency to handle
// long-running functions.
func noLimit(urls []string) {
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			result, err := request(u)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("result:", result)
		}(url)
	}

	wg.Wait()
}

// limited improves the `noLimit` implementation by introducing a limit to
// the number of concurrent operations that are processed at once. It makes
// use of a buffered channel to determine if a job can proceed or if it needs
// to block until a free slot is available.
//
// In short, trying to write a new value to a buffered channel that is "full"
// is a blocking operation. Thus, by attempting to write to the buffered
// channel at the start of the goroutine, we can "pause" the goroutine if
// there are no available slot in the channel until such time as the channel
// accepts the value.
func limited(urls []string) {
	wg := sync.WaitGroup{}
	limitChan := make(chan bool, 2)

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			// We use a deferred func here so that we don't need to issue the
			// channel read at every possible exit point of the goroutine.
			defer func() { <-limitChan }()
			limitChan <- true

			result, err := request(u)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println("result:", result)
		}(url)
	}

	wg.Wait()
}

// LimitedPool is implements an object to coordinate limited concurrency.
// Jobs can be queued up and will not be evaluated until the [LimitedPool.Run]
// method is invoked.
type LimitedPool struct {
	limitChan chan bool
	jobs      map[int]PoolJob

	Results chan string
	Errors  chan error
}

func (lp *LimitedPool) AddJob(job PoolJob) {
	l := len(lp.jobs)
	lp.jobs[l+1] = job
}

func (lp *LimitedPool) IsEmpty() bool {
	return len(lp.jobs) == 0
}

func (lp *LimitedPool) RemoveJob(index int) {
	delete(lp.jobs, index)
}

func (lp *LimitedPool) Run() {
	for i := 0; i < len(lp.jobs); i += 1 {
		job := lp.jobs[i]
		go func(j PoolJob, i int) {
			defer func() { <-lp.limitChan }()
			defer func() { lp.RemoveJob(i) }()
			lp.limitChan <- true

			result, err := j()
			if err != nil {
				lp.Errors <- err
				return
			}
			lp.Results <- result
		}(job, i)
	}
}

type PoolJob func() (string, error)

func NewLimitedPool(limit int, jobs ...PoolJob) *LimitedPool {
	pool := &LimitedPool{
		limitChan: make(chan bool, limit),
		jobs:      make(map[int]PoolJob),

		Results: make(chan string),
		Errors:  make(chan error),
	}

	for i := 0; i < len(jobs); i += 1 {
		pool.jobs[i] = jobs[i]
	}

	return pool
}
