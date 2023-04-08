package main

import "fmt"

// --program--
// 1. something that runs / executes code (go run somefile.go)
// => main() -> entry  function of the program
// => package main

// --package / library / module --
// 2. something that someone can import into their program or library
// -> package <whatever>

// go build -o myapp
func main() {
	user := User{
		username: "James",
		age:      getNumber(),
	}
	fmt.Println("the user is:", user)
}
