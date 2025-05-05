package models

type UserSignUp struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Message  string `json:"message"`
}

type UserLogin struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type UserLoginReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
