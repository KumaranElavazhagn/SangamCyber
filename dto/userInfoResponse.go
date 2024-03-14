package dto

type InsertUserResponse struct {
	UserId int `json:"userId"`
}

type AuthUserResponse struct {
	UserId int    `json:"userId"`
	Result string `json:"result"`
}
