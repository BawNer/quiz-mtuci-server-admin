package entity

type User struct {
	ID         int    `json:"id"`
	GroupID    int    `json:"-"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	PassText   string `json:"-"`
	NumberZach string `json:"numberZach,omitempty"`
	IsStudent  int    `json:"isStudent"`
	Token      string `json:"-"`
	Group      *Group `json:"group"`
}
