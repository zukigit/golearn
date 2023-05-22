package package1

import "fmt"

func MultipleReturns() {
	fmt.Println()
	fmt.Println("multiple returns")
	fmt.Println(helloMultiple())

	_, b := helloMultiple1()

	fmt.Println(b)

	num, s, in := helloMultiple()
	fmt.Println(num, s, in)
}

func helloMultiple() (int, string, int) {
	return 0, "zuki", 11
}

func helloMultiple1() (int, string) {
	return 0, "zuki"
}
