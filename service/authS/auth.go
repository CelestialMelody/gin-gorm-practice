package authS

import "gin-gorm-practice/models/blogAuth"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() error {
	return blogAuth.CheckAuth(a.Username, a.Password)
}
