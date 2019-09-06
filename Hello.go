package main
import "fmt"

var name = "Atiqur Rahman"

func main () {
	var name = "Atiq"
	fmt.Println("Hello -",name)
	foo()
}

func string foo () {
	var x = name + "How Are you"
	fmt.Printf("%T %v\n", x, x)
}