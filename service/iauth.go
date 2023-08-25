package service

type AuthInterface interface {
	Auth(username string, password string) bool
}
