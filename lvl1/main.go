package main

import (
	"fmt"

	"zuki.com/test/package1"
)

func main() {
	package1.HelloWorldTest("Hello World")
	Devide()

	package1.ValuesTest("Values Test")
	Devide()

	package1.VariableTest("Variable Test")
	Devide()

	package1.ConstantTest("Constant Test")
	Devide()

	package1.ForLoop("For Loop Test")
	Devide()

	package1.Arrays()
	Devide()

	package1.IfElse()
	Devide()

	package1.GoSwitch()
	Devide()

	package1.Arrays()
	Devide()

	package1.Slice()
	Devide()

	package1.Maps()
	Devide()

	package1.GoRange()
	Devide()

	package1.Functions("zuki")
	Devide()

	package1.MultipleReturns()
	Devide()

	package1.VariadicFunction(1, 2, 3, 4, 5, 6)
	Devide()

	package1.Closures()
	Devide()

	package1.Recursion()
	Devide()

	package1.Pointers()
	Devide()

	package1.StringsRunes()
	Devide()

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
	Devide()

	package1.Generics(1234)
}

func doJob(interf package1.StudentFunction) {
	fmt.Println(interf.GreetStudent())
	fmt.Println(interf.GoodByeStudent())
}

func Devide() {
	fmt.Println("-----------------------")
}
