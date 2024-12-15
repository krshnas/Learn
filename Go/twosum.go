package main

// import (
// 	"fmt"
// 	"time"
// )

// func twoSum(nums []int, target int) (int, int) {
// 	hash := make(map[int]int)
// 	for i, num := range nums {
// 		complement := target - num
// 		if j, ok := hash[complement]; ok {
// 			return j, i
// 		}
// 		hash[num] = i
// 	}
// 	return -1, -1 // not found
// }

// func main() {
// 	nums := []int{2, 7, 11, 15}
// 	target := 9

// 	// Record the start time
// 	start := time.Now()

// 	// Call the twoSum function
// 	i, j := twoSum(nums, target)

// 	// Record the end time and calculate the duration
// 	duration := time.Since(start)

// 	// Output the result
// 	fmt.Printf("Indices: %d, %d\n", i, j)
// 	fmt.Printf("Execution Time: %v\n", duration)
// }
