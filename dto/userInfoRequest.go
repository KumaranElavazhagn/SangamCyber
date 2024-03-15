package dto

type UserInfoRequest struct {
	UserName string `json:"userName" validate:"required,min=4,max=20"`
	Password string `json:"password" validate:"required,min=6"`
	EmailID  string `json:"emailId" validate:"required,email"`
}

type AuthUserInfoRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
