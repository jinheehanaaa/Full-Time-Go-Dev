package main

import "fmt"

type Player struct {
	HP int
}

// if player is not a pointer we are adjusting the copy of the player
// not the actual player.
func TakeDamage(player Player, amount int) {

	player.HP -= amount
	fmt.Println("player is taking damage. New HP ->", player.HP)
}

func main() {
	player := Player{
		HP: 100,
	}
	TakeDamage(player, 10)
	fmt.Printf("%+v\n", player)
}
