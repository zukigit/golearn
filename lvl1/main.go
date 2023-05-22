package main

import (
	"fmt"

	"zuki.com/test/package1"
)

func main() {
	package1.Arrays()
	package1.IfElse()
	package1.GoSwitch()
	package1.Arrays()
	package1.Slice()
	package1.Maps()
	package1.GoRange()
	package1.Functions("zuki")
	package1.MultipleReturns()
	package1.VariadicFunction(1, 2, 3, 4, 5, 6)
	package1.Closures()
	package1.Recursion()
	package1.Pointers()
	package1.StringsRunes()
	s1 := package1.NewStudent("zuki", 22, "somewhere")
	s2 := package1.Student{
		Name:    "misaki",
		Age:     22,
		Address: "somewhere",
	}
	fmt.Println()
	fmt.Println("interface")
	doJob(s2)
	doJob(s1)
	fmt.Println()
	fmt.Println("Struct Embedding")

	b1 := package1.Book{
		Student: package1.Student{
			Name: "miii",
		},
		Name: "some book name",
	}
	fmt.Println(b1.GreetStudent())
	b1.RentBook()
	b1.Test()
	b1.Student.Test()
	fmt.Println(b1)

	package1.Generics(1234)
}

func doJob(interf package1.StudentFunction) {
	fmt.Println(interf.GreetStudent())
	fmt.Println(interf.GoodByeStudent())
}
