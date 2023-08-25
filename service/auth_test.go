package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	type args struct {
		username string
		password string
	}
	test_cases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "in-valid test case",
			args: args{username: "test1", password: "test"},
			want: false,
		},
		{
			name: "valid test case",
			args: args{username: "test", password: "test"},
			want: true,
		},
	}

	for _, test := range test_cases {
		t.Run(test.name, func(t *testing.T) {
			aut_svc := AuthServcie{}
			got := aut_svc.Auth(test.args.username, test.args.password)
			assert.Equal(t, test.want, got)

		})
	}
}
