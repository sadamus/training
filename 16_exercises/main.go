package main

// func half(x int) (int, bool) {
// 	return x / 2, x%2 == 0
// }

func main() {

	// half := func(n int) (float64, bool) {
	// 	return float64(n) / 2, n&2 == 0
	// }

	// h, even := half(8)

	// fmt.Println(h, even)

	greatest := func(sf ...int) int {
		var largest int
		for _, v := range sf {
			if v > largest {
				largest = v
			}
		}
		return largest
	}

	println("The greatest is: ", greatest(1000, 56, 4096, 4, 7, 39))

}
