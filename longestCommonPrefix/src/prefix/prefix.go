package prefix

import "strings"


func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i:=1; i < len(strs); i++ {
		for strings.HasPrefix(strs[i], prefix) != true {
			prefix = prefix[0:len(prefix) - 1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}