package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrNoWorkers           = errors.New("errors no worker")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n <= 0 {
		return ErrNoWorkers
	}

	taskCh := make(chan Task)
	var errsCount int32
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for t := range taskCh {
				err := t()
				if err != nil {
					atomic.AddInt32(&errsCount, 1)
				}
			}
		}()
	}

	hasError := checkResult(taskCh, tasks, &errsCount, m)

	wg.Wait()

	if hasError {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func checkResult(ch chan Task, tasks []Task, errCount *int32, m int) bool {
	for _, t := range tasks {
		if atomic.LoadInt32(errCount) == int32(m) {
			close(ch)
			return true
		}
		ch <- t
	}
	close(ch)
	return false
}
