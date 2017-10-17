package main

import "fmt"

func main() {
	//Creating a map through make
	var myGreeting = make(map[string]string)
	//Set some values
	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Bonjour"
	myGreeting["Mike"] = "Gutan Tag"
	fmt.Println(myGreeting)

	//Creating a map through a composite literal
	yourGreeting := map[string]string{
		"Josh":   "What's up??",
		"Harold": "Good Day, Sir.",
	}
	//Still add new key:value through assignment
	yourGreeting["Peter"] = "How do you do?"
	//Reassigning a value
	yourGreeting["Peter"] = "May I help you?"
	//Actually, let's delete it
	delete(yourGreeting, "Peter")

	fmt.Println(yourGreeting)
	fmt.Println(len(yourGreeting))

}
