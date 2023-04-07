package main

import "fmt"

var (
	floatVar32 float32 = 0.1
	floatVar64 float64 = 0.1
	name       string  = "Foo"
	intVar32   int32   = 1       // 4 bytes
	intVar64   int64   = 4546456 // 8 bytes
	intVar     int     = -1
	uintVar    uint    = 1
	uint32Var  uint32  = 1
	uint64Var  uint64  = 10
	uint8Var   uint8   = 0x1 // same as byte
	byteVar    byte    = 0x2 // same as uint8
	runeVar    rune    = 'a' // string
)

type Player struct {
	name        string
	health      int
	attackPower float64
}

func (player Player) getHealth() int {
	return player.health
}

func main() {
	numbers := []int{}             // empty slice
	otherNumbers := make([]int, 0) // empty slice
	fmt.Println(numbers)
	fmt.Println(otherNumbers)
}
