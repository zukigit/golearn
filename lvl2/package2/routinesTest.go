package package2

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Loop(message string) {
	for i := 0; i < 11; i++ {
		fmt.Println(message, i)
	}
	wg.Done()
}

func RoutinesTest(message string) {
	fmt.Println(message)

	wg.Add(3)

	go Loop("hi loop")

	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg, i)
		}
		wg.Done()
	}("going")

	wg.Wait()
	fmt.Println("done")
}
