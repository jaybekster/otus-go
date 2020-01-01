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
	taskResultCh := make(chan error)
	hibernateCh := make(chan struct{})
	canExitCh := make(chan struct{})

	for i := 0; i < limit; i++ {
		go func(taskCh <- chan Worker, taskResultCh chan <- error, hibernateCh <- chan struct{}, canExitCh chan <- struct{}) {
			defer func() {
				canExitCh <- struct{}{}
			}()

			for {
				select {
				case task := <- taskCh:
					taskResultCh <- task()
				case <- hibernateCh:
					return
				}
			}
		}(taskCh, taskResultCh, hibernateCh, canExitCh)
	}

	go func() {
		for _, task := range tasks {
			select {
			case taskCh <- task:
			case <- hibernateCh:
				return
			}
		}
	}()

	func() {
		var counter int
		var errors int

		for {
			err := <- taskResultCh
			counter++

			fmt.Println("Complteted tasks:", counter)
			fmt.Println("Result: ", err)

			if err != nil {
				errors++
			}

			fmt.Println("counter", counter)

			if counter == len(tasks) || errors == maxErrors {
				close(hibernateCh)
				break;
			}
		}
	}()

	for i := 0; i < limit; i++ {
		select {
		case <- canExitCh:
			continue
		case <- taskResultCh:
			i = i - 1
		}
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
