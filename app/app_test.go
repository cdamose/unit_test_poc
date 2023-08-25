package app

import (
	"fmt"
	"testing"

	mocks "github.com/cdamose/unit_test_poc/mocks/repomocks"
	"github.com/stretchr/testify/assert"
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
			name: "Possitive flow auth",
			args: Args{username: "test1", password: "test"},
			want: false,
		},
	}
	for _, test := range testcase {
		t.Run(test.name, func(t *testing.T) {
			auth_svc_mock := &mocks.AuthInterface{}

			auth_svc_mock.On("Auth", "test", "test").
				Return(true).Once()

			auth_svc_mock.On("Auth", "test1", "test").
				Return(false).Once()

			//mock auth svc
			// mock_auth_svc := nil
			app := NewAuthApp(auth_svc_mock)
			got := app.Auth(test.args.username, test.args.password)
			fmt.Println(got)
			assert.Equal(t, test.want, got)
		})
	}
}
