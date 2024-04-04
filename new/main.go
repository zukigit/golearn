package main

import (
	"fmt"
	newpkg "zukidono/main/newPkg"
)

var age int

func main() {
	fmt.Println(newpkg.SayHi("zuki"))
	fmt.Println("age is", age)
}
