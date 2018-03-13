package main

import "fmt"
import "median"

func main ()  {
	A := []int{1, 2}
	B := []int{3, 4}
	mid := median.FindMedianSortedArrays(A, B)
	fmt.Println(mid)
}
