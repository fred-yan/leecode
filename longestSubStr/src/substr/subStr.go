package substr

//import "fmt"

func LengthOfLongestSubstring(s string) int {
	maxLength := 0
	maxs := 0
	maxe := 0
	p := 0
	q := 0

	charMap := make(map[uint8] int, 200)
	for i:=0; i<200; i++ {
		charMap[uint8(i)] = -1
	}

	length := len(s)
	for i:=0; i < length; i++ {
		q = i
		if charMap[s[i]] == -1 {
			charMap[s[i]] = i
		} else {
			index, ok := charMap[s[i]]
			if ok {
				if (q - p) > maxLength {
					maxs = p
					maxe = q
					maxLength = q - p
				}
				//fmt.Println(maxLength, index, p, q)
				p = index + 1
				charMap[s[index]] = i
			}
		}
	}

	if (q - p) > maxLength {
		return q - p
	} else {
		return maxe - maxs
	}
}