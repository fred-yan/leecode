package main

import "fmt"
import "addtwo"

func main() {
	t1 := addtwo.New()
	t2 := addtwo.New()
	t1.Insert(1,2)
	t1.Insert(2,8)
	t1.Insert(3,9)
	t2.Insert(1,3)
	t2.Insert(2,7)
	t2.Insert(3,9)

	t3 := addtwo.AddTwoNumbers(t1, t2)

	fmt.Println("print t1 and t2")
	t1.Traverse()
	t2.Traverse()
	t3.Traverse()
}
