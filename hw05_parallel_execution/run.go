package hw05parallelexecution

import (
	"errors"
	"sync"
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

	guardCh := make(chan struct{}, n)
	doneCh := make(chan bool, n)
	var s, e int
	var wg sync.WaitGroup

	for _, task := range tasks {
		guardCh <- struct{}{}

		wg.Add(1)
		go func(Task) {
			defer wg.Done()
			defer func() {
				<-guardCh
			}()

			err := task()
			doneCh <- err == nil
		}(task)

		err := checkResults(doneCh, &n, &m, &s, &e)
		if err != nil {
			return err
		}
	}

	wg.Wait()
	close(guardCh)
	close(doneCh)

	return nil
}

func checkResults(doneCh <-chan bool, n *int, m *int, s *int, e *int) (err error) {
	done := <-doneCh

	if done {
		*s++
	} else {
		*e++
		if *e >= *m || (*e+*s > *n+*m) {
			err = ErrErrorsLimitExceeded
		}
	}

	return err
}
