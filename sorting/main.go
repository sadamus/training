package main

import (
	"fmt"
	"sort"
)

/*

type people []string

studyGroup := people{"Zeno", "John", "Al", "Jenny"}

s := []string{"Zeno", "John", "Al", "Jenny"}

n := []int{7, 4, 8, 2, 9, 19, 12, 32,3}

*/

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	fmt.Println("Welcome to the game!")

	// sort a slice by name
	s := []string{"Zeno", "John", "Al", "Jenny"}

	sort.Strings(s)
	fmt.Println("Study Group: ", s)

	// sort a slice of int
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println("ints: ", n)

	// sort a slice of type people
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

}
