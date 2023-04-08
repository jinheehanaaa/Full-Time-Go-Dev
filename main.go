package main

import "fmt"

func main() {
	name := "foo42343"

	switch name {
	case "Alice":
		fmt.Println("The name Alice")
	case "Bob":
		fmt.Println("The name Bob")
	default:
		fmt.Println("the name is default =>", name)
	}
}
