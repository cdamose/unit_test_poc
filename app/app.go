package app

import "github.com/cdamose/unit_test_poc/service"

type AuthApp struct {
	authSvc service.AuthInterface
}

func NewAuthApp(authSvc service.AuthInterface) AuthInterface {
	return &AuthApp{authSvc: authSvc}
}

func (app *AuthApp) Auth(username string, password string) bool {
	return false
}
