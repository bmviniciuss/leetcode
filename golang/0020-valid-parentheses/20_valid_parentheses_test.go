package problem20

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
