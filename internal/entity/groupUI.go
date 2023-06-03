package entity

type GroupsResponse struct {
	Success     bool     `json:"success"`
	Description string   `json:"description"`
	Groups      []*Group `json:"groups"`
}
