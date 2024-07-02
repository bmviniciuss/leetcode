package problem14

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			strs: []string{"dog"},
			want: "dog",
		},
		{
			strs: []string{""},
			want: "",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			assert.Equal(t, tt.want, got)
		})
	}
}
