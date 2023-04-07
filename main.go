package main

import "fmt"

// weapon type
// axe
// sword
// wooden stick
// knife

func getDamage(weaponType string) int {
	switch weaponType {
	case "axe":
		return 100
	case "sword":
		return 90
	case "woodenStick":
		return 1
	case "knife":
		return 40
	default:
		panic("weapon does not exist")
	}
}

func main() {
	fmt.Println("damage of the given weapon:", getDamage("dd"))
}
