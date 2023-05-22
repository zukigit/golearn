package package1

import "fmt"

func ForLoop() {
	fmt.Println()
	fmt.Println("forLoop")
	var i, j = 0, 0

	for i < 5 {
		fmt.Println("it's for loop", i)
		i++
	}

	for {
		if j == 5 {
			break
		} else {
			fmt.Println("it's second for loop", j)
			j++
		}
	}

	for k := 0; k < 6; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println("it's not devisable by 2")
	}
}
