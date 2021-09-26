package main

import "fmt"

// lengthOfLongestSubstring lengthOfLongestSubstring
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var (
		max, num int
		strMap   = make(map[int32]struct{}, 0)
	)

	for l, byt := range s {
		if _, ok := strMap[byt]; ok {
			if max < num {
				max = num
			}
			num = 1
			continue
		}

		strMap[byt] = struct{}{}
		num++

		if len(s) == l+1 {
			if max < num {
				max = num
			}
			return max
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("au"))
}
