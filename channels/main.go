package main

import (
	"fmt"
)

//const p = "fmt.Println"

// func main() {
//p := "fmt.Println"
// c := make(chan int)
// go func() {
// 	for i := 0; i < 10; i++ {
// 		c <- i
// 	}
// 	close(c)
// }()

//for i := 0; i < 10; i++ {
// for n := range c {
// 	// fmt.Println(<-c)
// 	fmt.Println(n)
// }
// //p("End of Program....")

func main() {
	c := factorial(4)
	for n := range c {
		fmt.Println(n)
	}
}
func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
