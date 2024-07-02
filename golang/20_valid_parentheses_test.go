package golang

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stack struct {
	values []rune
}

func NewStack() *Stack {
	return &Stack{
		values: []rune{},
	}
}

func (s *Stack) Insert(val rune) {
	s.values = append(s.values, val)
}

func (s *Stack) Peek() (*rune, bool) {
	if len(s.values) == 0 {
		return new(rune), false
	}
	idx := len(s.values) - 1
	val := &s.values[idx]
	return val, true
}

func (s *Stack) Pop() (*rune, bool) {
	if len(s.values) == 0 {
		return new(rune), false
	}
	idx := len(s.values) - 1
	val := s.values[idx]
	s.values = s.values[:idx]
	return &val, true
}

func (s *Stack) Len() int {
	return len(s.values)
}

var (
	openPara     = rune('(')
	openCurly    = rune('{')
	openBrackt   = rune('[')
	closePara    = rune(')')
	closeCurly   = rune('}')
	closeBrackts = rune(']')
	closingMap   = map[rune]rune{
		closePara:    openPara,
		closeCurly:   openCurly,
		closeBrackts: openBrackt,
	}
)

func isOpeningRune(r rune) bool {
	return r == openPara || r == openCurly || r == openBrackt
}

func isClosingRune(r rune) bool {
	return r == closePara || r == closeCurly || r == closeBrackts
}

func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}
	if l%2 != 0 {
		return false
	}

	stack := NewStack()
	for _, letterRune := range s {
		if isOpeningRune(letterRune) {
			stack.Insert(letterRune)
			continue
		}
		if isClosingRune(letterRune) {
			complement := closingMap[letterRune]
			top, ok := stack.Peek()
			if !ok {
				return false // got closing but stack is empty
			}
			if *top != complement {
				return false // does not match with stack top
			}
			stack.Pop()
		}
	}
	return stack.Len() == 0
}

func Test_isValid(t *testing.T) {
	tests := []struct {
		arg  string
		want bool
	}{
		{arg: "()", want: true},
		{arg: "()[]{}", want: true},
		{arg: "(]", want: false},
		{arg: "", want: true},
		{arg: "(", want: false},
		{arg: "[", want: false},
		{arg: "{", want: false},
		{arg: "}", want: false},
		{arg: "]", want: false},
		{arg: ")", want: false},
		{arg: "([{[{[{(())}]}]}])", want: true},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, tt.arg), func(t *testing.T) {
			assert.Equal(t, tt.want, isValid(tt.arg))
		})
	}
}

func Test_Stack(t *testing.T) {
	t.Run("Should be able to add items to stack and remove the last two ones", func(t *testing.T) {
		stack := NewStack()
		assert.Equal(t, 0, stack.Len())

		stack.Insert(rune('1'))
		assert.Equal(t, 1, stack.Len())
		stack.Insert(rune('2'))
		assert.Equal(t, 2, stack.Len())
		stack.Insert(rune('3'))
		assert.Equal(t, 3, stack.Len())

		r, ok := stack.Pop()
		assert.Equal(t, rune('3'), *r)
		assert.True(t, ok)
		assert.Equal(t, 2, stack.Len())

		r, ok = stack.Pop()
		assert.Equal(t, rune('2'), *r)
		assert.True(t, ok)
		assert.Equal(t, 1, stack.Len())

		r, ok = stack.Pop()
		assert.Equal(t, rune('1'), *r)
		assert.True(t, ok)
		assert.Equal(t, 0, stack.Len())

		r, ok = stack.Pop()
		assert.Equal(t, new(rune), r)
		assert.False(t, ok)
	})
}
