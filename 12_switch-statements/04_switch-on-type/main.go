package main

import (
	"fmt"
)

type Contact struct {
	greeting string
	name     string
}

type Warrior struct {
	hpBonus     int16
	pwrBonus    int16
	spdDeficit  int16
	battlecry   string
	prfWeapType string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Printf("%v is an integer \n", x)
	case string:
		fmt.Printf("%v is a string \n", x)
	case Contact:
		fmt.Printf("%v is an Contact\n", x)
	case Warrior:
		fmt.Printf("%v is an Warrior. A WAAAARRRRRIIOOORRR!\n" ,x)
	default:
		fmt.Printf("%v is of an unknown type \n", x)
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Shawn Stevens")
	SwitchOnType("Warrior")
	//Create Warrior
	var garen = Warrior{20, 50, 15, "WAAARRRRRRIIIIOOORRR", "BattleAxe"}
	//Create Contact
	var heyJoe = Contact{"Heeyyy, Joe.", "Joe Hendrix"}
	//No case made for condition of type bool, defaults out
	var falseVar = false
	//Switch of struct Contact
	SwitchOnType(heyJoe)
	//Switch a made struct Warrior
	SwitchOnType(garen)
	//Switch an attritube of a Warrior struct
	SwitchOnType(garen.battlecry)
	//Switch on bool type with no case condition
	SwitchOnType(falseVar)
}
