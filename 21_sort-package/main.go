package main

import (
	"fmt"
	"sort"
)

type people struct {
	name []string
}

func main() {
	var studyGroup = people{
			[]string{"Zeno", "John", "Al", "Jenny"},
	}

	fmt.Println(studyGroup.name)
	fmt.Println("Are the strings sorted? ", sort.StringsAreSorted(studyGroup.name))

	sort.Strings(studyGroup.name)

	fmt.Println(studyGroup.name)
	fmt.Println("Are the strings sorted? ", sort.StringsAreSorted(studyGroup.name))

	var reversed sort.StringSlice = studyGroup.name

	sort.Sort(sort.Reverse(reversed))
	fmt.Println("Reversing..\n", reversed)

	// Might as well

	n := []int{100,64,38,997,16,1337,52,46,90,29,44,9,18,32}
	fmt.Println(n)
	//Create an IntSlice out of the slice of int
	//Sort it,
	sort.Ints(n)
	fmt.Println(n)
	//then reverse sort it.
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}

