package main

import "fmt"

type Player struct {
	HP int
}

// if player is not a pointer we are adjusting the copy of the player
// not the actual player.
func TakeDamage(player *Player, amount int) {

	player.HP -= amount
	fmt.Println("player is taking damage. New HP ->", player.HP)
}

// Function receiver
func (p *Player) TakeDamage(amount int) {
	p.HP -= amount
	fmt.Println("PLAYER IS TAKING DAMAGE. nEW hp ->", p.HP)
}

func main() {
	player := &Player{
		HP: 100,
	}
	//TakeDamage(player, 10) // Function
	player.TakeDamage(10) // Method
	fmt.Printf("%+v\n", player)
}
