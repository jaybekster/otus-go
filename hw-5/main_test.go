package main

import (
	"testing"
	"fmt"
	"time"
	"errors"
	"reflect"
)

func TestOneWorkerOneError(t *testing.T) {
	expectResut := []int{1, 2, 3, 4}
	testResult := make([]int, 0)

	task1 := func() error {
		time.Sleep(6 * time.Second)

		testResult = append(testResult, 1)

		return nil
	}

	task2 := func() error {
		time.Sleep(3 * time.Second)

		testResult = append(testResult, 2)

		return nil
	}

	task3 := func() error {
		time.Sleep(1 * time.Second)

		testResult = append(testResult, 3)

		return errors.New("error from task3")
	}

	task4 := func() error {
		time.Sleep(1 * time.Second)

		testResult = append(testResult, 4)

		return nil
	}

	task5 := func() error {
		time.Sleep(1 * time.Second)

		testResult = append(testResult, 5)

		return nil
	}

	Run([]Worker{task1, task2, task3, task4, task5}, 1, 1)

	isEqual := reflect.DeepEqual(expectResut, testResult)

	if !isEqual {
		fmt.Println(expectResut, testResult)
		t.Errorf("Not correct")
	}
}

func  TestExitOnCompleteAll(t *testing.T) {
	expectResut := []int{1, 2, 3, 4}
	testResult := make([]int, 0)

	task1 := func() error {
		time.Sleep(1 * time.Second)

		testResult = append(testResult, 1)

		return nil
	}

	task2 := func() error {
		time.Sleep(2 * time.Second)

		testResult = append(testResult, 2)

		return nil
	}

	task3 := func() error {
		time.Sleep(3 * time.Second)

		testResult = append(testResult, 3)

		return nil
	}

	task4 := func() error {
		time.Sleep(4 * time.Second)

		testResult = append(testResult, 4)

		return nil
	}


	Run([]Worker{task1, task2, task3, task4}, 2, 1)

	isEqual := reflect.DeepEqual(expectResut, testResult)

	if !isEqual {
		fmt.Println(expectResut, testResult)
		t.Errorf("Not correct")
	}
}

func TestTwoWorkersTwoErrors(t *testing.T) {
	expectResut := []int{1, 2, 3, 4}
	testResult := make([]int, 0)

	task1 := func() error {
		time.Sleep(1 * time.Second)

		testResult = append(testResult, 1)

		return errors.New("error 1")
	}

	task2 := func() error {
		time.Sleep(2 * time.Second)

		testResult = append(testResult, 2)

		return errors.New("error 2")
	}

	task3 := func() error {
		time.Sleep(3 * time.Second)

		testResult = append(testResult, 3)

		return nil
	}

	task4 := func() error {
		time.Sleep(4 * time.Second)

		testResult = append(testResult, 4)

		return nil
	}

	task5 := func() error {
		time.Sleep(1 * time.Second)

		testResult = append(testResult, 5)

		return nil
	}

	task6 := func() error {
		time.Sleep(2 * time.Second)

		testResult = append(testResult, 6)

		return nil
	}


	Run([]Worker{task1, task2, task3, task4, task5, task6}, 2, 2)

	isEqual := reflect.DeepEqual(expectResut, testResult)

	fmt.Println(expectResut, testResult)

	if !isEqual {
		fmt.Println(expectResut, testResult)
		t.Errorf("Not correct")
	}
}
