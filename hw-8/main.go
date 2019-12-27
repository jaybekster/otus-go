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
	resultCh := make(chan error)
	closeCh := make(chan struct{})
	closedCh := make(chan struct{}, limit)

	isHibernating := false

	// allocate the tickets:
	for i := 0; i < limit; i++ {
		go func(taskCh <- chan Worker, resultCh chan <- error, closeCh <- chan struct{}, closedCh chan <- struct{}) {
			defer func() {
				fmt.Println("send to closed")
				closedCh <- struct{}{}
			}()

			for {

				select {
				
				case task, ok := <- taskCh:
					if !ok {
						return
					}

					value :=  task()

					if !isHibernating {
						resultCh <- value
					}
				
				case <- closeCh:
					return
	
				default:
	
				}
			}

		}(taskCh, resultCh, closeCh, closedCh)
	}


	go func() {
		defer func() {
			fmt.Println("end iterating")
		}()

		for _, task := range tasks {
			select {
			case taskCh <- task:
			case <- closeCh:
				return
			}
		}
	}()


	var counter int
	var errors int

	for {
		err := <- resultCh

		counter += 1

		if err != nil {
			errors += 1
		}

		if counter == len(tasks) || errors == maxErrors {
			isHibernating = true
			close(closeCh)
			break;
		}
	}

	fmt.Println("wait")

	for i := 0; i < limit; i +=1 {
		<- closedCh
		fmt.Println("closed success")
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

	Run([]Worker{task1, task2, task3, task4, task5, task6}, 5, 1)

	fmt.Println("number of goroutines: ", runtime.NumGoroutine())
}
