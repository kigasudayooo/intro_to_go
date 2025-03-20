package main

import (
	"fmt"

	"awesomeProject/mylib"
	"awesomeProject/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.say()
	under.Hello()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
}
