package main

import "fmt"

func grettings() string {
	a := "Hello, Go enthusiastic."
	return a
}

func name() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return "Hello! Welcome back, " + name
}


func main() {
	r := grettings()
	fmt.Println(r)

	n := name()
	fmt.Println(n)
}
