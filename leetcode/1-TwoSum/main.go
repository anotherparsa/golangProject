package main

import "fmt"

func main() {
	numbers1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println(twoSumFirst(numbers1, target1))
	fmt.Println(twoSumFSecond(numbers1, target1))

	numbers2 := []int{3, 2, 4}
	target2 := 6

	fmt.Println(twoSumFirst(numbers2, target2))
	fmt.Println(twoSumFSecond(numbers2, target2))

	numbers3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSumFirst(numbers3, target3))
	fmt.Println(twoSumFSecond(numbers3, target3))
}

func twoSumFirst(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumFSecond(numbers []int, target int) []int {
	alreadySeenNumbers := map[int]int{}
	for i := 0; i < len(numbers); i++ {
		x := target - numbers[i]
		if value, ok := alreadySeenNumbers[x]; ok {
			return []int{value, i}
		}
		alreadySeenNumbers[numbers[i]] = i
	}
	return nil
}
