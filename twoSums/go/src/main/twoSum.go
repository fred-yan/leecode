package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var tMap map[int]int
	tMap = make(map[int]int)
	arr := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		tMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		value, ok := tMap[complement]
		if ok && (value != i) {
			arr[0] = i
			arr[1] = value
			break
		}
	}
	return arr

}

func main() {
	res := make([]int, 2)
	nums := [4]int{2, 7, 11, 15}
	target := 9
	var numslice []int = nums[:]
	res = twoSum(numslice, target)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

}
