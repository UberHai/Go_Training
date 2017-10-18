package main

import "fmt"

type Player struct {
	Name		string
	Health		int
	Class 		string
}

func (p Player) Attack() string {
	fmt.Println(p.Name, "attacks with all his might")
	return p.Name + "attacks with all his might"
}

type Attacker interface {
	Attack() string
}

func DoAttack(a Attacker) {
	_ = a.Attack()
}

func main() {
	jimmy := Player{"Jimbob91", 99, "Warrior"}
	DoAttack(jimmy)
}