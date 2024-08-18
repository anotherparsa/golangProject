package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	numbers1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSum(numbers1, target1))

	numbers2 := []int{3, 2, 4}
	target2 := 6

	fmt.Println(twoSum(numbers2, target2))

	numbers3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(numbers3, target3))
}
