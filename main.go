package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	userID, err := fetchUserID()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the response took %v -> %+v\n", time.Since(start), userID)
}

func fetchUserID() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	type result struct {
		userId string
		err    error
	}

	resultch := make(chan result, 1)

	go func() {
		res, err := thirdpartyHTTPCall()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	// Done()
	// 1. -> the context timeout is exceeded
	// 2. -> the context has been manually canceled -> Cancel()
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err
	}

}

func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 101)
	return "user id 1", nil
}
