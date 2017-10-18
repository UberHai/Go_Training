package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type SecretAgent struct {
	Person
	LicenseToKill bool
}

//func [receiver] name([parameter...]) return
//This function is received by the person type. Returns a string
func (p Person) fullName() string {
	return p.First + " " + p.Last
}
//This function goes to the SecretAgent type
func (s SecretAgent) confirmLicense() string {
	if s.LicenseToKill == true {
		return "Why yes, I am licensed to kill."
	}
	return "No, I am not licensed to kill"
}

func main() {
	var p3 SecretAgent
	p3.First = "James"
	p3.Last = "Bond"
	p3.Age = 24
	p3.LicenseToKill = true

	p4 := SecretAgent{
		Person: Person{
			First: "Miss",
			Last: "MoneyPenny",
			Age: 20,

		},
		LicenseToKill: false,
	}

	fmt.Println(p3.fullName())
	fmt.Println(p4.fullName())
}
