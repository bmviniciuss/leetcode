package problem13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
