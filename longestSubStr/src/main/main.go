package main

import "substr"
import "fmt"

func main ()  {
	str1 := "helloworld"
	str2 := "abcabcbb"
	str3 := "bbbbb"
	str4 := "pwwkew"
	num1 := substr.LengthOfLongestSubstring(str1)
	num2 := substr.LengthOfLongestSubstring(str2)
	num3 := substr.LengthOfLongestSubstring(str3)
	num4 := substr.LengthOfLongestSubstring(str4)
	fmt.Println(num1, num2, num3, num4)
}

