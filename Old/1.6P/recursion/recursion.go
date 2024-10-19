package main

import "fmt"

func main() {
	fact := factorial(4)
	fmt.Println(fact)
}

func factorial(num int) int {
	// Recursion
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
	// result := 1
	// for i := 1; i <= num; i++ {
	// 	result *= i
	// }
	// return result
}
