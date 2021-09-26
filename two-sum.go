package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, 0)
	for key, val := range nums {
		if _, ok := numsMap[target-val]; ok {
			return []int{numsMap[target-val], key}
		}

		numsMap[val] = key
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}
