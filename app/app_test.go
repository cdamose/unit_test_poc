package app

import (
	"testing"
)

func TestAuth(t *testing.T) {
	type Args struct {
		username string
		password string
	}
	testcase := []struct {
		name string
		args Args
		want bool
	}{
		{
			name: "Possitive flow auth",
			args: Args{username: "test", password: "test"},
			want: true,
		},
		{
			name: "Negative flow auth",
			args: Args{username: "test1", password: "test"},
			want: false,
		},
	}
	for _, test := range testcase {
		t.Run(test.name, func(t *testing.T) {
			//mock auth svc
			// mock_auth_svc := nil
			// app := NewAuthApp(mock_auth_svc)
			// got := app.Auth(test.args.username, test.args.password)
			// assert.Equal(t, test.want, got)
		})
	}
}
