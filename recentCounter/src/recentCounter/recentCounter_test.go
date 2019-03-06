package recentCounter

import (
	"testing"
	"fmt"
)

func TestRecentCounter_Ping(t *testing.T) {
	obj := Constructor()
	pingTime := []int{1, 100, 3001, 3002, 7000}
	for _, t := range pingTime {
		param1 := obj.Ping(t)
		fmt.Println(param1)
	}
}