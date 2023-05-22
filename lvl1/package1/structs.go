package package1

type Student struct {
	Name    string
	Age     int
	Address string
}

func NewStudent(name string, age int, address string) Student {
	return Student{Name: name, Age: age, Address: address}
}
