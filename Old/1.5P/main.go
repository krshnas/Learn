package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Rae"
	userNames = append(userNames, "Krishna")
	userNames = append(userNames, "Singh")

	fmt.Println(userNames)

	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 2.0
	courseRatings["react"] = 4.0
	courseRatings["angular"] = 4.5
	// courseRatings["node"] = 4.1

	// fmt.Println(courseRatings)
	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("Index:", index, "\tValue:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key, "\tValue:", value)
	}
}
