package main

import "fmt"
import "prefix"

func main ()  {
	var strs = []string{"leecode", "leets", "leesafsd"}
	result := prefix.LongestCommonPrefix(strs)
	fmt.Println(result)
}