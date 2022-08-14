package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	intMap := map[int]int{}

	for i, n := range nums {
		if val, ok := intMap[target-n]; ok {
			return []int{val, i}
		} else {
			intMap[n] = i
		}
	}

	return nil
}
