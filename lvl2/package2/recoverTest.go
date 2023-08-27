package package2

import "fmt"

func RecoverTest(name string) {
	fmt.Println(name)

	defer func() {
		if recvr := recover(); recvr != nil {
			fmt.Println("Recovered, problem is\n", recvr)
		}
	}()

	panic("unknown problem")
	fmt.Println("after panic")
}