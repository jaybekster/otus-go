package main

import (
	"errors"
	"fmt"
	"time"
	"runtime"
)

type Worker = func() error


func Allocate(start <- chan Worker) {
	go func(start <- chan Worker) {
		
	}(start)
}

func Run(tasks []Worker, limit, maxErrors int) {
	taskCh := make(chan Worker)
	resultCh := make(chan error)
	closeCh := make(chan struct{})
	closedCh := make(chan struct{}, limit)

	// allocate the tickets:
	for i := 0; i < limit; i++ {
		go func(taskCh <- chan Worker, resultCh chan <- error, closeCh <- chan struct{}, closedCh chan <- struct{}) {
			defer func() {
				fmt.Println("send to closed")
				closedCh <- struct{}{}
			}()

			for {
				select {
				
				case task := <- taskCh:
					resultCh <- task()
				
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
			taskCh <- task;
		}
	}()


	var counter int

	for {
		err := <- resultCh

		fmt.Println(err)
		counter += 1

		if counter == len(tasks) {
			fmt.Println("end")
			close(closeCh)
			break;
		}
	}

	fmt.Println("wait")

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

	// task4 = func() error {
	// 	time.Sleep(4 * time.Second)

	// 	return nil
	// }

	// task5 := func() error {
	// 	time.Sleep(4 * time.Second)

	// 	return nil
	// }

	// task6 := func() error {
	// 	time.Sleep(4 * time.Second)

	// 	return nil
	// }

	Run([]Worker{task1, task2, task3}, 2, 10)

	fmt.Println("number of goroutines: ", runtime.NumGoroutine())
}
