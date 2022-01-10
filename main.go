package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

type Task struct {
	Name           string
	Type           string
	Directorate    string
	ExecuteTime    time.Time
	ExternalResult string
}

func GenerateTask(i int) Task {

	var t Task

	switch i % 3 {
	case 0:
		t = Task{
			Name:        fmt.Sprintf("Mr. Dif-%d", i),
			Type:        "Difficult",
			Directorate: "Cyber",
			ExecuteTime: time.Now(),
		}
	case 1:
		t = Task{
			Name:        fmt.Sprintf("Mr. Izzy-%d", i),
			Type:        "Izzy Pizzy",
			Directorate: "Marketing",
			ExecuteTime: time.Now(),
		}
	case 2:
		t = Task{
			Name:        fmt.Sprintf("Mr. NoMik-%d", i),
			Type:        "No Mikir",
			Directorate: "Hedonia",
			ExecuteTime: time.Now(),
		}
	}

	return t

}

func main() {

	WithoutSemaphore()
	SemaphoreImplementation1()
	SemaphoreImplementation2()
	SemaphoreImplementation3()

}

func WithoutSemaphore() {

	for i := 0; i < 5000; i++ {
		fmt.Printf("Number of Go Routine : %d\n", runtime.NumGoroutine())
		go func(i int) {
			t := GenerateTask(i)
			fmt.Printf("i : %d -> t.Name : %s\n", i, t.Name)
		}(i)
	}

}

func SemaphoreImplementation1() {

	sem := make(chan bool, 10)

	for i := 0; i < 5000; i++ {
		fmt.Printf("Number of Go Routine : %d\n", runtime.NumGoroutine())

		// if semaphore full , it will block
		sem <- true
		go func(i int) {
			defer func() {
				<-sem // release
			}()

			t := GenerateTask(i)
			fmt.Printf("i : %d -> t.Name : %s\n", i, t.Name)
		}(i)
	}

}

func SemaphoreImplementation2() {

	sem := semaphore.NewWeighted(10)

	for i := 0; i < 5000; i++ {
		fmt.Printf("for 1 Number of Go Routine : %d\n", runtime.NumGoroutine())

		// if semaphore full , it will block
		if err := sem.Acquire(context.Background(), 1); err != nil {
			log.Fatal(err)
		}

		go func(i int) {
			defer func() {
				sem.Release(1) // release
			}()

			t := GenerateTask(i)
			fmt.Printf("i : %d -> t.Name : %s\n", i, t.Name)
		}(i)
	}
}

func SemaphoreImplementation3() {

	sem := semaphore.NewWeighted(10)

	rc := make(chan error)
	var j int

	j++
	go func() {

		defer func() {
			rc <- nil
		}()

		for i := 0; i < 5000; i++ {
			fmt.Printf("for 1 Number of Go Routine : %d\n", runtime.NumGoroutine())

			// if semaphore full , it will block
			if err := sem.Acquire(context.Background(), 1); err != nil {
				log.Fatal(err)
			}

			go func(i int) {
				defer func() {
					sem.Release(1) // release
				}()

				t := GenerateTask(i)
				fmt.Printf("i : %d -> t.Name : %s\n", i, t.Name)
			}(i)
		}
	}()

	j++
	go func() {
		defer func() {
			rc <- nil
		}()

		for i := 0; i < 5000; i++ {
			fmt.Printf("for 2 Number of Go Routine : %d\n", runtime.NumGoroutine())

			// if semaphore full , it will block
			if err := sem.Acquire(context.Background(), 1); err != nil {
				log.Fatal(err)
			}

			go func(i int) {
				defer func() {
					sem.Release(1) // release
				}()

				t := GenerateTask(i)
				fmt.Printf("i : %d -> t.Name : %s\n", i, t.Name)
			}(i)
		}
	}()

	j++
	go func() {

		defer func() {
			rc <- nil
		}()

		for i := 0; i < 5000; i++ {
			fmt.Printf("for 3 Number of Go Routine : %d\n", runtime.NumGoroutine())

			// if semaphore full , it will block
			if err := sem.Acquire(context.Background(), 1); err != nil {
				log.Fatal(err)
			}

			go func(i int) {
				defer func() {
					sem.Release(1) // release
				}()

				t := GenerateTask(i)
				fmt.Printf("i : %d -> t.Name : %s\n", i, t.Name)
			}(i)
		}
	}()

	for i := 0; i < j; i++ {
		<-rc
	}
}
