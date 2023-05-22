package package1

import (
	"fmt"
)

func (s Student) GreetStudent() string {
	return "hello " + s.Name
}

func (s Student) GoodByeStudent() string {
	return "goodbye " + s.Name
}

func (book Book) RentBook() {
	fmt.Println(book.Name, "is rented")
}

func (student Student) Test() {
	fmt.Println("tested for student", student.Name)
}

func (book Book) Test() {
	fmt.Println("tested for book", book.Name)
}