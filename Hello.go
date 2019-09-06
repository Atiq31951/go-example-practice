package main
import "fmt"

var name = "Atiqur Rahman"

func main () {
	var name = "Atiq"
	fmt.Println("Hello -",name)
	foo()
}

func foo () {
	var x = name + "How Are you"
	fmt.Printf("Type of x is => %T, and Value is ==> %s\n", x, x)
}