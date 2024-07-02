package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stack struct {
	values []int
}

func (s *Stack) Append(val int) {
	s.values = append(s.values, val)
}

func (s *Stack) Pop() int {
	val := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val
}

func (s *Stack) Len() int {
	return len(s.values)
}

func (s *Stack) Peek() int {
	return s.values[len(s.values)-1]
}

func longestValidParentheses(s string) int {
	maxLen := 0
	stack := &Stack{}
	stack.Append(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Append(i)
		} else {
			stack.Pop()
			if stack.Len() == 0 {
				stack.Append(i)
			} else {
				maxLen = max(maxLen, i-stack.Peek())
			}
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
