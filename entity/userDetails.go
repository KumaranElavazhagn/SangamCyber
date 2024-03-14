package entity

type UserID struct {
	UserID int `db:"user_id"`
}

type UserInfoResponse struct {
	UsernameInfo []UserID `json:"username_info"`
	EmailInfo    []UserID `json:"email_info"`
}

type AuthEntityResponse struct {
	UserID   int    `db:"user_id"`
	Password string `db:"password"`
}
