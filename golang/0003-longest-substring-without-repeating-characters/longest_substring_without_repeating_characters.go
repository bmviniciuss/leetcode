package problem0003

func lengthOfLongestSubstring(s string) int {
	var (
		left   = 0
		seqMax = 0
		set    = map[byte]struct{}{}
	)
	for right := 0; right < len(s); right++ {
		for {
			if _, ok := set[s[right]]; !ok {
				break
			}
			delete(set, s[left])
			left++
		}
		set[s[right]] = struct{}{}
		seqMax = max(seqMax, right-left+1)
	}
	return seqMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	if len(s) == 1 {
// 		return 1
// 	}
// 	var (
// 		left      = 0
// 		seqMax    = 0
// 		letterSet = map[string]struct{}{}
// 	)
// 	for {
// 		if left >= len(s) {
// 			break
// 		}
// 		for right := left; right < len(s); right++ {
// 			letter := string(s[right])
// 			_, ok := letterSet[letter]
// 			if !ok {
// 				letterSet[letter] = struct{}{}
// 				continue
// 			}
// 			left++
// 			seqMax = max(seqMax, len(letterSet))
// 			letterSet = map[string]struct{}{}
// 			break
// 		}
// 	}

// 	return seqMax
// }
