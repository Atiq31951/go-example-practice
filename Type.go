package main
import ("fmt")

func main() {
	var s string = "Hello"
	var i int
	for i = 0; i < len(s); i++ {
		// if (s[i] == 'H') {
		// 	fmt.Println("Hello")
		// 	break
		// }
		if s[i] == 'H' {
			fmt.Println("Hello")
			break
		}
	}
}
