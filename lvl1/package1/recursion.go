package package1

import "fmt"

func Recursion() {
	fmt.Println()
	fmt.Println("recursion")

	var recurFunction func(i int) int

	recurFunction = func(i int) int {
		if i < 2 {
			return i
		} else {
			return recurFunction(i - 1)
		}
	}

	finalNum := recurFunction(17)
	fmt.Println("final number is", finalNum)
}
