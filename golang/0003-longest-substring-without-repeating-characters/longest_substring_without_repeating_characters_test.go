package problem0003

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		{arg: "abcabcbb", want: 3},
		{arg: "bbbbb", want: 1},
		{arg: "pwwkew", want: 3},
		{arg: "au", want: 2},
		{arg: "abcb", want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, tt.arg), func(t *testing.T) {
			assert.Equal(t, tt.want, lengthOfLongestSubstring(tt.arg))
		})
	}
}
