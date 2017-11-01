package main

import "fmt"

func main() {
	var tMap map[int]int
	tMap = make(map[int]int)

	nums := [4]int{2, 7, 11, 15}
	target := 9

	for i := 0; i < len(nums); i++ {
		tMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		value, ok := tMap[complement]
		if ok && (value != i) {
			fmt.Println(i, value)
			break
		} else {
			fmt.Println("Error")
		}
	}

}