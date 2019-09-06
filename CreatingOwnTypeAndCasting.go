package main
import "fmt"

type str string

var name str

func main () {
	var s str = "My name is atiq"
	fmt.Printf("Hello, %T\n", s)

	var s1 string

	fmt.Printf("Before typecasting type of s1 => %T\n", s)

	s1 = string(s)
	fmt.Printf("After casting s1 type =  %T\n", s1)

}