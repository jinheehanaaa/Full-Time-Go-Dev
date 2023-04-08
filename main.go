package main

import "fmt"

// Declaration of interface
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

// Changing to other DB
type PostgresNumberStore struct {
	// postgres values (db connection)
}

func (p PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6, 7, 8}, nil
}

func (p PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the Postgres storage")
	return nil
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
		numberStore: PostgresNumberStore{},
	}

	// Logic
	if err := apiServer.numberStore.Put(1); err != nil {
		panic(err)
	}

	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	// no error
	fmt.Println(numbers)
}
