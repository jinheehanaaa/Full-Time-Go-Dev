package main

import "fmt"

type Position struct {
	x, y int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

// Avoid duplication by struct embedding
type SpecialEntity struct {
	Entity
	specialRField float64
}

func main() {

	e := SpecialEntity{
		specialRField: 88.88,
		Entity: Entity{
			//name:    "my special entity",
			version: "1.0",
			//id:      "my special id",
			Position: Position{
				x: 100,
				y: 200,
			},
		},
	}
	e.id = "my special id"
	e.name = "foo"
	e.x = 350
	e.y = 500
	fmt.Printf("%+v\n", e.Position)
}
