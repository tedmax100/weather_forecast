package model

type UserLoginRequest struct {
	Name     string
	Password string
}

func (user *UserLoginRequest) IsValid() bool {
	return !(user.Name == "" && user.Password == "")
}

type UserLoginResponse struct {
	Token string
}
