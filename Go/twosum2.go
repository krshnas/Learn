package main

// import "go.uber.org/zap"

// func twoSum(target int, arr []int) (int, int) {
// 	sumHash := make(map[int]int)

// 	for k, num := range arr {
// 		diff := target - num
// 		if l, ok := sumHash[diff]; ok {
// 			return l, k
// 		}
// 		sumHash[num] = k
// 	}
// 	return -1, -1
// }

// func main() {
// 	logger := zap.Must(zap.NewProduction())
// 	defer logger.Sync()
// 	target := 10
// 	arr := []int{1, -11, 7, 12, 9}
// 	a, b := twoSum(target, arr)
// 	logger.Info("Two sum result", zap.Any("first num", a), zap.Any("second num", b))
// }
