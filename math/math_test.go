package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	test_cases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return added number",
			args: args{10, 20},
			want: 30,
		},
	}
	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			got := sum(test.args.a, test.args.b)
			assert.Equal(t, test.want, got)
		})
	}

}
