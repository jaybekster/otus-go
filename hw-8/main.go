package main

import (
	"errors"
	"fmt"
	"time"
	"runtime"
)

type Worker = func() error

func Run(tasks []Worker, limit, maxErrors int) {
	taskCh := make(chan Worker)
	resultCh := make(chan error, limit -1)
	closeCh := make(chan struct{})
	closedCh := make(chan struct{})

	for i := 0; i < limit; i++ {
		go func(taskCh <- chan Worker, resultCh chan <- error, closeCh <- chan struct{}, closedCh chan <- struct{}) {
			defer func() {
				closedCh <- struct{}{}
			}()


			for {
				select {
				case task := <- taskCh:
					resultCh <- task()
				case <- closeCh:
					return
				}
			}
		}(taskCh, resultCh, closeCh, closedCh)
	}

	func() {
		var counter int
		var errors int
		var inProgress int

		i := 0

		for ; i < limit; i++ {
			inProgress++
			taskCh <- tasks[i]
		}

		for {
			err := <- resultCh
			inProgress--
			counter++

			fmt.Println("Complteted tasks:", counter)
			fmt.Println("Result: ", err)

			if err != nil {
				errors++
			}

			if counter == len(tasks) || errors == maxErrors {
				close(closeCh)
				return
			} else if len(tasks) - counter - inProgress > 0 {
				inProgress++
				taskCh <- tasks[limit -1 + counter]
			}
		}
	}()

	for i := 0; i < limit; i +=1 {
		<- closedCh
	}
}

func main() {
	var task1, task2, task3 Worker;

	task1 = func() error {
		time.Sleep(6 * time.Second)

		return nil
	}

	task2 = func() error {
		time.Sleep(3 * time.Second)

		return nil
	}

	task3 = func() error {
		time.Sleep(1 * time.Second)

		return errors.New("error from task3")
	}

	task4 := func() error {
		time.Sleep(4 * time.Second)

		return nil
	}

	task5 := func() error {
		time.Sleep(4 * time.Second)

		return nil
	}

	task6 := func() error {
		time.Sleep(4 * time.Second)

		return nil
	}

	Run([]Worker{task1, task2, task3, task4, task5, task6}, 2, 1)

	fmt.Println("number of goroutines: ", runtime.NumGoroutine())
}
