package problem14

func longestCommonPrefix(strs []string) string {
	slen := len(strs)
	if slen == 0 {
		return ""
	}
	if slen == 1 {
		return strs[0]
	}

	var (
		min = minLen(strs)
		res = ""
	)
	if min == 0 {
		return res
	}
	for i := 0; i < min; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}
	return res
}

func minLen(strs []string) int {
	if len(strs) <= 0 {
		return 0
	}
	if len(strs) == 1 {
		return len(strs[0])
	}

	min := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}
	return min
}
