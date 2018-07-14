// app.go
package main

import "fmt"

// Reads a web page ==============================================

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	res, _ := http.Get("https://www.nasdaq.com/symbol/cldt/dividend-history")
// 	page, _ := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	fmt.Printf("%s", page)
// }

/*
 * Input a name from the keyboard ===================================
 */
// import "fmt"

// func main() {
// 	var name string
// 	fmt.Print("Please enter your name: ")
// 	fmt.Scan(&name)
// 	fmt.Println("Hello %s", name)
// }

/* Enter two numbers and divide them
 */

// import "fmt"

// func main() {
// 	var numOne int
// 	var numTwo int
// 	fmt.Print("Enter a large number: ")
// 	fmt.Scan(&numOne)
// 	fmt.Print("Enter a smaller number: ")
// 	fmt.Scan(&numTwo)
// 	fmt.Println(numOne, "/", numTwo, " = ", numOne/numTwo)
// }

/*
   Print all even numbers between 0 & 100
*/

var p = fmt.Println

func main() {
	s := "1 Year Target"
	p(s, ":", s[2])
	if s[1] == 32 {
		p("Yes")
	}
}
