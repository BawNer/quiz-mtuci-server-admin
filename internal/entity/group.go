package entity

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
}

type GroupViaJoin struct {
	GID   int    `json:"id"`
	GName string `json:"name"`
	GYear string `json:"year"`
}
