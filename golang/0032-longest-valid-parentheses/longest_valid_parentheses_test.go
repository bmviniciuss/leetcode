package problem32

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestValidParentheses(t *testing.T) {
	tests := []struct {
		arg  string
		want int
	}{
		{arg: "", want: 0},
		{arg: "(()", want: 2},
		{arg: ")()())", want: 4},
		{arg: "()(())", want: 6},
		{arg: "()(()", want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, tt.arg), func(t *testing.T) {
			assert.Equal(t, tt.want, longestValidParentheses(tt.arg))
		})
	}
}
