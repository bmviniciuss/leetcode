package golang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "III",
			arg:  "III",
			want: 3,
		},
		{
			name: "IV",
			arg:  "IV",
			want: 4,
		},
		{
			name: "LVIII",
			arg:  "LVIII",
			want: 58,
		},
		{
			name: "MCMXCIV",
			arg:  "MCMXCIV",
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.arg)
			assert.Equal(t, tt.want, got)
		})
	}
}
