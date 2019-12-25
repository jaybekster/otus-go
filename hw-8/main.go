package main

import (
	"errors"
	"fmt"
	"time"
)

type Worker = func() error

func ExecuteWorker(task chan Worker, result chan error) {
	go func() {
		for {
			task, ok := <- task

			if !ok {
				return
			}

			result <- task()
		}
	}()
}

func Run(tasks []Worker, limit, maxErrors int) {
	start := make(chan Worker, limit - 1)
	finish := make(chan error, limit)
	done := make(chan interface{})

	var counter int
	var errors int

	for i := 0; i < limit; i++ {
		ExecuteWorker(start, finish)
	}


	go func() {
		for _, task := range tasks {
			start <- task
		}
	}()

	go func() {
		for err := range finish {
			counter += 1
			if err != nil {
				fmt.Println(err)
				errors += 1
			}

			if errors == maxErrors || counter == len(tasks) {
				done <- nil
			}
		}
	}()

	<- done
	close(start)
	close(finish)
}

func main() {
	var task1, task2, task3, task4 Worker;

	task1 = func() error {
		time.Sleep(6 * time.Second)
		fmt.Println("Task 1")

		return errors.New("error from task1")
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

	Run([]Worker{task1, task2, task3, task4}, 2, 1)
}
