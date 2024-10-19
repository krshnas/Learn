package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumpup(1, 10, 15)

	anotherSum := sumpup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumpup(staringvalue int, numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
