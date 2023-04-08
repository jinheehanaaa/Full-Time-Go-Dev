package main

import (
	"fmt"
	"mygoproject/types"
	"mygoproject/util"
)

// User -> public access everywhere
// user -> private access BUT public ijn its own package

// go build -o myapp
func main() {
	user := types.User{
		Username: util.GetUsername(),
		Age:      util.GetAge(),
	}
	fmt.Printf("the user is: %+v", user)
}
