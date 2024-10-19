package lists

import "fmt"

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{45.70, 21.70, 99.99, 70.70}

// 	fmt.Println(prices)
// 	productNames[2] = "A carpet"
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// featuredPrices := prices[1:3] // start from 2nd pos to 4th pos, but skipping the 4th element
// 	// featuredPrices := prices[:3] // start from beginning to 4th pos, but skipping the 4th element
// 	featuredPrices := prices[1:]            // start from 2nd pos to end, include last element too // highest bound could be last element + 1
// 	highlightedPrices := featuredPrices[:1] // will include only first element
// 	featuredPrices[0] = 199.99
// 	fmt.Println(featuredPrices)
// 	fmt.Println(highlightedPrices)
// 	// fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))       // len 3 cap 3
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // len 1 cap 3

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)                              // will contain all the elements from featuredPrices, as it has capicity (more on right), it won't do that for left side, whatever missed is forever missed
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // len 3 cap 3
// }

func main() {
	prices := []float64{10.99, 8.99}

	fmt.Println(prices, len(prices), cap(prices))
	prices = append(prices, 5.99)
	fmt.Println(prices, len(prices), cap(prices))

	discountPrices := []float64{121.78, 80.12, 90.12}
	prices = append(prices, discountPrices...) // ... (three dots) allowed to pull all the values from slice / array to append
	fmt.Println(prices)
}
