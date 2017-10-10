package main

import "fmt"

func main() {
	variable := "MyConditionIs15" //15 Characters long
	switch {
	case len(variable) == 10:
		fmt.Println("Your length was ten. You may pass.")
	case variable == "MyConditionIs12":
		fmt.Println("You probably won't see me")
	case len(variable)+len(variable) == 10:
		fmt.Print("Ahh. Five plus five equals ten. I see")
	case len(variable) == 15:
		fmt.Println("Ahh. Path chosen wisely my boy.")
	default:
		fmt.Println("I couldn't match any of your conditions. Try again")
	}
}
