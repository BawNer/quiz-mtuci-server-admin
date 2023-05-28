package entity

type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ResponseUserLogin struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}
