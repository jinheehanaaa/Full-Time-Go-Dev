package main

import "fmt"

// Declaration of interface
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

// Implementation of interface
type MongoDBNumberStore struct {
	// some values
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the mogoDB storage")
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	// no error
	fmt.Println(numbers)
}
