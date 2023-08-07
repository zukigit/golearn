package package2

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("jobs number", j)
		results <- j * 4
		time.Sleep(time.Second * 2)
	}
}

func WorkerPoolsTest(msg string) {
	fmt.Println(msg)

	numberOfJobs := 5
	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)

	for i := 0; i < 3; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < numberOfJobs; i++ {
		jobs <- i
	}

	for i := 0; i < numberOfJobs; i++ {
		res := <-results
		fmt.Println("return result for job", i, "is", res)
	}
}
