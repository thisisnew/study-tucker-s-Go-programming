package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕 난 %d살 %s라고해.", s.Age, s.Name)
}

func main() {
	student := Student{
		Name: "철수",
		Age:  12,
	}

	var stringer Stringer
	stringer = student

	fmt.Println(stringer.String())
}
