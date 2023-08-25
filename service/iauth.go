package service

type AuthInterface interface {
	auth(username string, password string) bool
}
