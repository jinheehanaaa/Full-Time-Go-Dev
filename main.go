package main

import (
	"fmt"
	"time"
)

func main() {
	resultch := make(chan string) // Unbuffered

	go func() {
		result := resultch
		fmt.Println(result)
	}()

	resultch <- "foo"
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}
