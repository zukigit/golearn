package package1

import "fmt"

func Functions(name string) {
	fmt.Println("")
	fmt.Println("functions")
	fmt.Println("hello", name)

	result := sum(10, 11)

	fmt.Println("sum of 10 and 11 is ", result)
}

func sum(num1, num2 int) int {
	return num1 + num2
}
