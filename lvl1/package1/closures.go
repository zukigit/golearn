package package1

import "fmt"

func Closures() {

	fmt.Println("closures")
	instance := createInstance()

	for {
		tem := instance()
		if tem == 5 {
			break
		} else {
			fmt.Println("instance value:", tem)
		}
	}
}

func createInstance() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
