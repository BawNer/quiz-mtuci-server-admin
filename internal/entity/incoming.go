package entity

type ErrorResponseUI struct {
	Description string `json:"description"`
	Code        string `json:"code"`
}

type PositiveResponseUI struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
}
