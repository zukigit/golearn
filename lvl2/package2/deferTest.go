package package2

import "fmt"

func DeferTest(name string) {
	fmt.Println(name)

	function1()
	defer function2()
	function3()
}

func function1() {
	fmt.Println("function1")
}

func function2() {
	fmt.Println("function2")
}

func function3() {
	fmt.Println("function3")
}