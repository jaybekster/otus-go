package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker = func() error

func ExecuteWorker(wg sync.WaitGroup, task Worker, finish chan <- interface{}) {
	finish <- task()

	wg.Done()
}

func Run(tasks []Worker, limit, maxErrors int) {
	var wg sync.WaitGroup

	start := make(chan Worker, limit)
	finish := make(chan interface{}, limit)
	done := make(chan interface{}, limit)
	var errors int

	for i := 0; i < limit; i++ {
		go func() {
			for {
				select {
				case task := <- start:
					ExecuteWorker(wg, task, finish)
				case <- done:
					return
				}
			}
		}()
	}

	for _, task := range tasks {
		select {
		case start <- task:
			wg.Add(1)
		case err := <- finish:
			if err != nil {
				errors += 1

			}
			
			if errors >= maxErrors {
				for i := 0; i < limit; i++ {
					done <- nil
				}

				break;
			}
		}
	}

	wg.Wait()
}

func main() {
	var task1, task2, task3, task4 Worker;

	task1 = func() error {
		time.Sleep(4 * time.Second)
		fmt.Println("Task 1")

		return nil;
	}

	task2 = func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("Task 2")

		return nil
	}

	task3 = func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("Task 3")

		return nil
	}

	task4 = func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("Task 4")

		return nil
	}

	Run([]Worker{task1, task2, task3, task4}, 3, 1)
}