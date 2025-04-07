package model

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Note    string `json:"note"`
	Completed bool `json:"completed"`
}
