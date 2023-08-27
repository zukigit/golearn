package package2

import (
	"fmt"
	"os"
)

func PanicTest(name string) {
	fmt.Println(name)

	_, err := os.Create("/root/Documents")
	if(err != nil) {
		panic(err)
	}
}