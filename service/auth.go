package service

type AuthServcie struct {
}

func NewAuthService() AuthInterface {
	return &AuthServcie{}
}
func (svc *AuthServcie) auth(username string, password string) bool {
	if username == "test" && password == "test" {
		return true
	} else {
		return false
	}

}
