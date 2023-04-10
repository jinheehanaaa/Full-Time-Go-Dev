package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	result, err := thirdpartyHTTPCall()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the response took %v -> %+v\n", time.Since(start), result)
}

func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 5000000)
	return "user id 1", nil
}
