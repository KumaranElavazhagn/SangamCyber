package dto

type UserInfoRequest struct {
	UserName    string `json:"userName" validate:"required,min=4,max=20"`
	Password    string `json:"password" validate:"required,min=6"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	EmailID     string `json:"emailId" validate:"required,email"`
	DateOfBirth string `json:"dateOfBirth" validate:"required,datetime=2006-01-02"`
}

type AuthUserInfoRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
