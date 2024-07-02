package problem13

var (
	numberVal = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

func romanToInt(s string) int {
	var (
		sum = 0
		i   = len(s) - 1
	)
	for i >= 0 {
		currentLetterValue := numberVal[string(s[i])]
		if i > 0 {
			lookAheadValue := numberVal[string(s[i-1])]
			if currentLetterValue > lookAheadValue {
				i--
				currentLetterValue -= lookAheadValue
			}
		}
		i--
		sum += currentLetterValue
	}
	return sum
}
